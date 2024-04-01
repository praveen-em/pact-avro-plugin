# Pact Avro Plugin

This plugin supports Avro encoded message paylod for the [Pact](http://docs.pact.io) framework.


## Repository Structure

```
├── go.mod                  # Go module                                     
├── main.go                 # Entrypoint for the application
├── plugin.go               # Stub gRPC methods (to be implemented)
├── configureinteraction/   # Implementation to configure pact interaction 
├── Makefile                # Build configuration
├── io_pact_plugin/         # Location of protobuf and gRPC definitions for Plugin Framework
├── log.go                  # Logging utility
├── pact-plugin.json        # Plugin configuration file
├── pact.go                 # Pact type definitions
├── server.go               # The gRPC server implementation
├── RELEASING.md            # Instructions on how to release 🚀
```

## Current State
The plugin in its current form can do the following
- Configure/build pact interaction with Avro encoded payload along with matching rules. Works for most of the primitive and complex data types.  
- Verify interaction. Supports only exact match at the moment.

**To do:**

Contributions are welcome! 

- Still some work required to translate the following Pact DSL logical types to appropriate avro format. This can be acheived by using information from avro schema or adding optional config (indicating avro types) for these logical types and updating Matching Rule Definition Grammar (.g4) accordingly. 

| Pact DSL Logical Type       | Avro                                  |
| --------------------------- | ------------------------------------- |
| integer                     | int, long                             |
| decimal                     | float, double, decimal (logical type) |
| number (integer or decimal) | int, long, float, double, decimal     |

- Matching rules path for eachValue $(reference) needs to be fixed. 
- Verify interaction based on Matching Rules. This will be easier to implement if FFI is already available for this.  
- Implement Generators. 

## Plugin usage examples
Refer to following example projects where this plugin is being used. Please note that these example projects are being constantly updated for various experimentation purposes, so they are not in a robust state yet.   

- [consumer java kafka example](https://github.com/praveen-em/example-consumer-java-kafka-avro)
- [provider java kafka example](https://github.com/praveen-em/example-provider-java-kafka-avro)

## Install
```
pact-plugin-cli -y install https://github.com/perodem/pact-avro-plugin/releases/tag/v0.0.1
```

## Developing plugin
### Local Development
The following command will build the plugin, and install into the correct plugin directory for local development:
 ```
 make install_local
 ```
You can then reference your plugin in local tests to try it out.

### Supported targets
This code base automatically create artifacts for the following OS/Architecture combinations:

| OS      | Architecture | Supported |
| ------- | ------------ | --------- |
| OSX     | x86_64       | ✅         |
| OSX     | arm          | ✅         |
| Linux   | x86_64       | ✅         |
| Linux   | arm          | ✅         |
| Windows | x86_64       | ✅         |
| Windows | arm          | ✅         |

### Publish plugin
Follow the steps in [Releasing](./RELEASING.md) to publish a new version of the Plugin. 