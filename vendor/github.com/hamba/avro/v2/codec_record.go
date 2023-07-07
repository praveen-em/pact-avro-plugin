package avro

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func createDecoderOfRecord(cfg *frozenConfig, schema Schema, typ reflect2.Type) ValDecoder {
	switch typ.Kind() {
	case reflect.Struct:
		return decoderOfStruct(cfg, schema, typ)

	case reflect.Map:
		if typ.(reflect2.MapType).Key().Kind() != reflect.String ||
			typ.(reflect2.MapType).Elem().Kind() != reflect.Interface {
			break
		}
		return decoderOfRecord(cfg, schema, typ)

	case reflect.Ptr:
		return decoderOfPtr(cfg, schema, typ)

	case reflect.Interface:
		if ifaceType, ok := typ.(*reflect2.UnsafeIFaceType); ok {
			return &recordIfaceDecoder{schema: schema, valType: ifaceType}
		}
	}

	return &errorDecoder{err: fmt.Errorf("avro: %s is unsupported for avro %s", typ.String(), schema.Type())}
}

func createEncoderOfRecord(cfg *frozenConfig, schema Schema, typ reflect2.Type) ValEncoder {
	switch typ.Kind() {
	case reflect.Struct:
		return encoderOfStruct(cfg, schema, typ)

	case reflect.Map:
		if typ.(reflect2.MapType).Key().Kind() != reflect.String ||
			typ.(reflect2.MapType).Elem().Kind() != reflect.Interface {
			break
		}
		return encoderOfRecord(cfg, schema, typ)

	case reflect.Ptr:
		return encoderOfPtr(cfg, schema, typ)
	}

	return &errorEncoder{err: fmt.Errorf("avro: %s is unsupported for avro %s", typ.String(), schema.Type())}
}

func decoderOfStruct(cfg *frozenConfig, schema Schema, typ reflect2.Type) ValDecoder {
	rec := schema.(*RecordSchema)
	structDesc := describeStruct(cfg.getTagKey(), typ)

	fields := make([]*structFieldDecoder, 0, len(rec.Fields()))
	for _, field := range rec.Fields() {
		sf := structDesc.Fields.Get(field.Name())
		if sf == nil {
			for _, alias := range field.Aliases() {
				sf = structDesc.Fields.Get(alias)
				if sf != nil {
					break
				}
			}
		}

		// Skip field if it doesnt exist
		if sf == nil {
			fields = append(fields, &structFieldDecoder{
				decoder: createSkipDecoder(field.Type()),
			})
			continue
		}

		dec := decoderOfType(cfg, field.Type(), sf.Field[len(sf.Field)-1].Type())
		fields = append(fields, &structFieldDecoder{
			field:   sf.Field,
			decoder: dec,
		})
	}

	return &structDecoder{typ: typ, fields: fields}
}

type structFieldDecoder struct {
	field   []*reflect2.UnsafeStructField
	decoder ValDecoder
}

type structDecoder struct {
	typ    reflect2.Type
	fields []*structFieldDecoder
}

func (d *structDecoder) Decode(ptr unsafe.Pointer, r *Reader) {
	for _, field := range d.fields {
		// Skip case
		if field.field == nil {
			field.decoder.Decode(nil, r)
			continue
		}

		fieldPtr := ptr
		for i, f := range field.field {
			fieldPtr = f.UnsafeGet(fieldPtr)

			if i == len(field.field)-1 {
				break
			}

			if f.Type().Kind() == reflect.Ptr {
				if *((*unsafe.Pointer)(fieldPtr)) == nil {
					newPtr := f.Type().(*reflect2.UnsafePtrType).Elem().UnsafeNew()
					*((*unsafe.Pointer)(fieldPtr)) = newPtr
				}

				fieldPtr = *((*unsafe.Pointer)(fieldPtr))
			}
		}
		field.decoder.Decode(fieldPtr, r)

		if r.Error != nil && !errors.Is(r.Error, io.EOF) {
			for _, f := range field.field {
				r.Error = fmt.Errorf("%s: %w", f.Name(), r.Error)
				return
			}
		}
	}
}

func encoderOfStruct(cfg *frozenConfig, schema Schema, typ reflect2.Type) ValEncoder {
	rec := schema.(*RecordSchema)
	structDesc := describeStruct(cfg.getTagKey(), typ)

	fields := make([]*structFieldEncoder, 0, len(rec.Fields()))
	for _, field := range rec.Fields() {
		sf := structDesc.Fields.Get(field.Name())
		if sf != nil {
			fields = append(fields, &structFieldEncoder{
				field:   sf.Field,
				encoder: encoderOfType(cfg, field.Type(), sf.Field[len(sf.Field)-1].Type()),
			})
			continue
		}

		if !field.HasDefault() {
			// In all other cases, this is a required field
			err := fmt.Errorf("avro: record %s is missing required field %q", rec.FullName(), field.Name())
			return &errorEncoder{err: err}
		}

		def := field.Default()
		if field.Default() == nil {
			if field.Type().Type() == Null {
				// We write nothing in a Null case, just skip it
				continue
			}

			if field.Type().Type() == Union && field.Type().(*UnionSchema).Nullable() {
				defaultType := reflect2.TypeOf(&def)
				fields = append(fields, &structFieldEncoder{
					defaultPtr: reflect2.PtrOf(&def),
					encoder:    encoderOfPtrUnion(cfg, field.Type(), defaultType),
				})
				continue
			}
		}

		defaultType := reflect2.TypeOf(def)
		fields = append(fields, &structFieldEncoder{
			defaultPtr: reflect2.PtrOf(def),
			encoder:    encoderOfType(cfg, field.Type(), defaultType),
		})
	}
	return &structEncoder{typ: typ, fields: fields}
}

type structFieldEncoder struct {
	field      []*reflect2.UnsafeStructField
	defaultPtr unsafe.Pointer
	encoder    ValEncoder
}

type structEncoder struct {
	typ    reflect2.Type
	fields []*structFieldEncoder
}

func (e *structEncoder) Encode(ptr unsafe.Pointer, w *Writer) {
	for _, field := range e.fields {
		// Default case
		if field.field == nil {
			field.encoder.Encode(field.defaultPtr, w)
			continue
		}

		fieldPtr := ptr
		for i, f := range field.field {
			fieldPtr = f.UnsafeGet(fieldPtr)

			if i == len(field.field)-1 {
				break
			}

			if f.Type().Kind() == reflect.Ptr {
				if *((*unsafe.Pointer)(fieldPtr)) == nil {
					w.Error = fmt.Errorf("embedded field %q is nil", f.Name())
					return
				}

				fieldPtr = *((*unsafe.Pointer)(fieldPtr))
			}
		}
		field.encoder.Encode(fieldPtr, w)

		if w.Error != nil && !errors.Is(w.Error, io.EOF) {
			for _, f := range field.field {
				w.Error = fmt.Errorf("%s: %w", f.Name(), w.Error)
				return
			}
		}
	}
}

func decoderOfRecord(cfg *frozenConfig, schema Schema, typ reflect2.Type) ValDecoder {
	rec := schema.(*RecordSchema)
	mapType := typ.(*reflect2.UnsafeMapType)

	fields := make([]recordMapDecoderField, len(rec.Fields()))
	for i, field := range rec.Fields() {
		fields[i] = recordMapDecoderField{
			name:    field.Name(),
			decoder: decoderOfType(cfg, field.Type(), mapType.Elem()),
		}
	}

	return &recordMapDecoder{
		mapType:  mapType,
		elemType: mapType.Elem(),
		fields:   fields,
	}
}

type recordMapDecoderField struct {
	name    string
	decoder ValDecoder
}

type recordMapDecoder struct {
	mapType  *reflect2.UnsafeMapType
	elemType reflect2.Type
	fields   []recordMapDecoderField
}

func (d *recordMapDecoder) Decode(ptr unsafe.Pointer, r *Reader) {
	if d.mapType.UnsafeIsNil(ptr) {
		d.mapType.UnsafeSet(ptr, d.mapType.UnsafeMakeMap(0))
	}

	for _, field := range d.fields {
		elem := d.elemType.UnsafeNew()
		field.decoder.Decode(elem, r)

		d.mapType.UnsafeSetIndex(ptr, reflect2.PtrOf(field), elem)
	}

	if r.Error != nil && !errors.Is(r.Error, io.EOF) {
		r.Error = fmt.Errorf("%v: %w", d.mapType, r.Error)
	}
}

func encoderOfRecord(cfg *frozenConfig, schema Schema, typ reflect2.Type) ValEncoder {
	rec := schema.(*RecordSchema)
	mapType := typ.(*reflect2.UnsafeMapType)

	fields := make([]mapEncoderField, len(rec.Fields()))
	for i, field := range rec.Fields() {
		fields[i] = mapEncoderField{
			name:    field.Name(),
			hasDef:  field.HasDefault(),
			def:     field.Default(),
			encoder: encoderOfType(cfg, field.Type(), mapType.Elem()),
		}

		if field.HasDefault() {
			switch {
			case field.Type().Type() == Union:
				union := field.Type().(*UnionSchema)
				fields[i].def = map[string]any{
					string(union.Types()[0].Type()): field.Default(),
				}
			case field.Default() == nil:
				continue
			}

			defaultType := reflect2.TypeOf(fields[i].def)
			fields[i].defEncoder = encoderOfType(cfg, field.Type(), defaultType)
			if defaultType.LikePtr() {
				fields[i].defEncoder = &onePtrEncoder{fields[i].defEncoder}
			}
		}
	}

	return &recordMapEncoder{
		mapType: mapType,
		fields:  fields,
	}
}

type mapEncoderField struct {
	name       string
	hasDef     bool
	def        any
	defEncoder ValEncoder
	encoder    ValEncoder
}

type recordMapEncoder struct {
	mapType *reflect2.UnsafeMapType
	fields  []mapEncoderField
}

func (e *recordMapEncoder) Encode(ptr unsafe.Pointer, w *Writer) {
	for _, field := range e.fields {
		valPtr := e.mapType.UnsafeGetIndex(ptr, reflect2.PtrOf(field))
		if valPtr == nil {
			// Missing required field
			if !field.hasDef {
				w.Error = fmt.Errorf("avro: missing required field %s", field.name)
				return
			}

			// Null default
			if field.def == nil {
				continue
			}

			defPtr := reflect2.PtrOf(field.def)
			field.defEncoder.Encode(defPtr, w)
			continue
		}

		field.encoder.Encode(valPtr, w)

		if w.Error != nil && !errors.Is(w.Error, io.EOF) {
			w.Error = fmt.Errorf("%s: %w", field.name, w.Error)
			return
		}
	}
}

type recordIfaceDecoder struct {
	schema  Schema
	valType *reflect2.UnsafeIFaceType
}

func (d *recordIfaceDecoder) Decode(ptr unsafe.Pointer, r *Reader) {
	obj := d.valType.UnsafeIndirect(ptr)
	if reflect2.IsNil(obj) {
		r.ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}

	r.ReadVal(d.schema, obj)
}

type structDescriptor struct {
	Type   reflect2.Type
	Fields structFields
}

type structFields []*structField

func (sf structFields) Get(name string) *structField {
	for _, f := range sf {
		if f.Name == name {
			return f
		}
	}

	return nil
}

type structField struct {
	Name  string
	Field []*reflect2.UnsafeStructField

	anon *reflect2.UnsafeStructType
}

func describeStruct(tagKey string, typ reflect2.Type) *structDescriptor {
	structType := typ.(*reflect2.UnsafeStructType)
	fields := structFields{}

	var curr []structField
	next := []structField{{anon: structType}}

	visited := map[uintptr]bool{}

	for len(next) > 0 {
		curr, next = next, curr[:0]

		for _, f := range curr {
			rtype := f.anon.RType()
			if visited[f.anon.RType()] {
				continue
			}
			visited[rtype] = true

			for i := 0; i < f.anon.NumField(); i++ {
				field := f.anon.Field(i).(*reflect2.UnsafeStructField)
				isUnexported := field.PkgPath() != ""

				chain := make([]*reflect2.UnsafeStructField, len(f.Field)+1)
				copy(chain, f.Field)
				chain[len(f.Field)] = field

				if field.Anonymous() {
					t := field.Type()
					if t.Kind() == reflect.Ptr {
						t = t.(*reflect2.UnsafePtrType).Elem()
					}
					if t.Kind() != reflect.Struct {
						continue
					}

					next = append(next, structField{Field: chain, anon: t.(*reflect2.UnsafeStructType)})
					continue
				}

				// Ignore unexported fields.
				if isUnexported {
					continue
				}

				fieldName := field.Name()
				if tag, ok := field.Tag().Lookup(tagKey); ok {
					fieldName = tag
				}

				fields = append(fields, &structField{
					Name:  fieldName,
					Field: chain,
				})
			}
		}
	}

	return &structDescriptor{
		Type:   structType,
		Fields: fields,
	}
}
