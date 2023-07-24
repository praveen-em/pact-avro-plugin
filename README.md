# Pact Avro Plugin

This plugin uses [Golang Plugin Template](https://github.com/pact-foundation/pact-plugin-template-golang) for the [Pact](http://docs.pact.io) framework. 

**Features:**

* Being built (WIP)

## Repository Structure

WIP - needs updating

```
├── go.mod            # Go module                              (✅ fill me in!)
├── main.go           # Entrypoint for the application
├── plugin.go         # Stub gRPC methods for you to implement (✅ fill me in!)
├── configuration.go  # Type definitions for your plugin's DSL (✅ fill me in!)
├── Makefile          # Build configuration                    (✅ fill me in!)
├── io_pact_plugin/   # Location of protobuf and gRPC definitions for Plugin Framework
├── log.go            # Logging utility
├── pact-plugin.json  # Plugin configuration file
├── pact.go           # Pact type definitions
├── server.go         # The gRPC server implementation
├── RELEASING.md      # Instructions on how to release 🚀
```

## Supported targets

This code base should automatically create artifacts for the following OS/Architecture combiations:

| OS      | Architecture | Supported |
| ------- | ------------ | --------- |
| OSX     | x86_64       | ✅         |
| OSX     | arm          | ✅         |
| Linux   | x86_64       | ✅         |
| Linux   | arm          | ✅         |
| Windows | x86_64       | ✅         |
| Windows | arm          | ✅         |