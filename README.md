# MOPS SDK

[![Go Version](https://img.shields.io/github/go-mod/go-version/totmicro/mops-sdk)](https://github.com/totmicro/mops-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/totmicro/mops-sdk.svg)](https://pkg.go.dev/github.com/totmicro/mops-sdk)
[![License](https://img.shields.io/github/license/totmicro/mops-sdk)](LICENSE)

The MOPS SDK provides the core interfaces and types for developing MOPS plugins. This SDK enables independent plugin development without requiring the full MOPS core system.

## Features

- 🔌 **Plugin Interface**: Standard interface for MOPS plugins
- ⚡ **Action System**: Rich action execution framework with parameter validation
- 🏗️ **Type System**: Comprehensive type definitions for MOPS operations
- 📊 **Version Management**: API compatibility checking between plugins and core
- 🛡️ **Validation**: Built-in parameter and configuration validation

## Quick Start

### Installing the SDK

```bash
go mod init my-mops-plugin
go get github.com/totmicro/mops-sdk@latest
```

### Creating a Simple Plugin

```go
package main

import (
    "fmt"
    "github.com/totmicro/mops-sdk/actions"
    "github.com/totmicro/mops-sdk/plugin"
    "github.com/totmicro/mops-sdk/types"
)

// Plugin implementation
type MyPlugin struct{}

// NewPlugin is the required export function
func NewPlugin() plugin.Plugin {
    return &MyPlugin{}
}

// Required plugin interface methods
func (p *MyPlugin) GetName() string        { return "my-plugin" }
func (p *MyPlugin) GetVersion() string     { return "1.0.0" }
func (p *MyPlugin) GetDescription() string { return "My awesome MOPS plugin" }
func (p *MyPlugin) GetAPIVersion() string  { return "1.0.0" }

func (p *MyPlugin) Initialize(config *plugin.PluginConfig) error {
    return nil
}

func (p *MyPlugin) GetProviders() []actions.DynamicProvider {
    return []actions.DynamicProvider{}
}

func (p *MyPlugin) GetExecutors() []actions.ActionExecutor {
    return []actions.ActionExecutor{&MyExecutor{}}
}

func (p *MyPlugin) GetFunctions() []actions.InteractiveFunction {
    return []actions.InteractiveFunction{}
}

func (p *MyPlugin) Cleanup() error {
    return nil
}

// Action executor implementation
type MyExecutor struct{}

func (e *MyExecutor) GetID() string {
    return "my-action"
}

func (e *MyExecutor) GetActions() []actions.ActionInfo {
    return []actions.ActionInfo{
        {
            ID:          "hello",
            Name:        "Hello World",
            Description: "Prints a hello message",
            Parameters: []actions.ParameterInfo{
                {
                    Name:        "name",
                    Type:        "string",
                    Description: "Name to greet",
                    Required:    false,
                    Default:     "World",
                },
            },
        },
    }
}

func (e *MyExecutor) Execute(actionID string, params map[string]interface{}) (*types.ActionResult, error) {
    name, ok := params["name"].(string)
    if !ok {
        name = "World"
    }
    
    return &types.ActionResult{
        Success: true,
        Message: fmt.Sprintf("Hello, %s!", name),
        Data:    map[string]interface{}{"greeting": name},
    }, nil
}

func (e *MyExecutor) Validate(actionID string, params map[string]interface{}) error {
    return nil
}

// Required for plugin compilation
func main() {}
```

### Building Your Plugin

```bash
# Build as shared library
go build -buildmode=plugin -o my-plugin.so .

# Or use the MOPS plugin build tools for comprehensive validation
```

## API Reference

### Core Interfaces

#### Plugin Interface
```go
type Plugin interface {
    GetName() string
    GetVersion() string
    GetDescription() string
    GetAPIVersion() string
    Initialize(config *PluginConfig) error
    GetProviders() []actions.DynamicProvider
    GetExecutors() []actions.ActionExecutor
    GetFunctions() []actions.InteractiveFunction
    Cleanup() error
}
```

#### ActionExecutor Interface
```go
type ActionExecutor interface {
    GetID() string
    GetActions() []ActionInfo
    Execute(actionID string, params map[string]interface{}) (*types.ActionResult, error)
    Validate(actionID string, params map[string]interface{}) error
}
```

### Package Overview

- **`types`**: Core data structures and result types
- **`actions`**: Action execution interfaces and metadata
- **`plugin`**: Plugin lifecycle and configuration interfaces  
- **`version`**: Version management and compatibility checking

## Examples

See the [examples](examples/) directory for complete plugin implementations:

- **Simple Plugin**: Basic plugin structure
- **Action Plugin**: Plugin with custom actions
- **Provider Plugin**: Plugin with dynamic menu providers
- **Complex Plugin**: Advanced plugin with multiple features

## Compatibility

The MOPS SDK follows semantic versioning:

- **1.0.x**: Stable API, backward compatible
- **1.x.x**: Minor features, backward compatible
- **x.x.x**: Major changes, may require plugin updates

Current API version: **1.0.0**

## Development

### Building the SDK

```bash
git clone https://github.com/totmicro/mops-sdk.git
cd mops-sdk
go mod tidy
go test ./...
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://pkg.go.dev/github.com/totmicro/mops-sdk)
- 🐛 [Issues](https://github.com/totmicro/mops-sdk/issues)
- 💬 [Discussions](https://github.com/totmicro/mops-sdk/discussions)
- 📧 [Support Email](mailto:support@totmicro.com)
