# Pact Avro Plugin

This plugin supports Avro encoded message paylod for the [Pact](http://docs.pact.io) framework. It is still being built, not ready yet. 

**Features:**

* Being built (WIP)

## Repository Structure

WIP - needs updating

```
├── go.mod            # Go module                              (✅ fill me in!)
├── main.go           # Entrypoint for the application
├── plugin.go         # Stub gRPC methods for you to implement (✅ fill me in!)
├── interaction_builder/*.go  # Code to configure/build interaction (✅ fill me in!)
├── Makefile          # Build configuration                    (✅ fill me in!)
├── io_pact_plugin/   # Location of protobuf and gRPC definitions for Plugin Framework
├── log.go            # Logging utility
├── pact-plugin.json  # Plugin configuration file
├── pact.go           # Pact type definitions
├── server.go         # The gRPC server implementation
├── RELEASING.md      # Instructions on how to release 🚀
```

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