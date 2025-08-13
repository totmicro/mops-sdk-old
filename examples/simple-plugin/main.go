package main

import (
	"fmt"
	"github.com/totmicro/mops-sdk/actions"
	"github.com/totmicro/mops-sdk/plugin"
	"github.com/totmicro/mops-sdk/types"
)

// Plugin implementation
type SimplePlugin struct {
	config *plugin.PluginConfig
}

// NewPlugin is the required export function
func NewPlugin() plugin.Plugin {
	return &SimplePlugin{}
}

// Required plugin interface methods
func (p *SimplePlugin) GetName() string {
	return "simple-plugin"
}

func (p *SimplePlugin) GetVersion() string {
	return "1.0.0"
}

func (p *SimplePlugin) GetDescription() string {
	return "A simple example MOPS plugin demonstrating basic functionality"
}

func (p *SimplePlugin) GetAPIVersion() string {
	return "1.0.0"
}

func (p *SimplePlugin) Initialize(config *plugin.PluginConfig) error {
	p.config = config
	fmt.Printf("Simple plugin initialized with config: %+v\n", config)
	return nil
}

func (p *SimplePlugin) GetProviders() []actions.DynamicProvider {
	// This plugin doesn't provide dynamic menu items
	return []actions.DynamicProvider{}
}

func (p *SimplePlugin) GetExecutors() []actions.ActionExecutor {
	return []actions.ActionExecutor{&SimpleExecutor{}}
}

func (p *SimplePlugin) GetFunctions() []actions.InteractiveFunction {
	// This plugin doesn't provide interactive functions
	return []actions.InteractiveFunction{}
}

func (p *SimplePlugin) Cleanup() error {
	fmt.Println("Simple plugin cleaned up")
	return nil
}

// Action executor implementation
type SimpleExecutor struct{}

func (e *SimpleExecutor) GetID() string {
	return "simple-actions"
}

func (e *SimpleExecutor) GetActions() []actions.ActionInfo {
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
		{
			ID:          "echo",
			Name:        "Echo Message",
			Description: "Echoes back the provided message",
			Parameters: []actions.ParameterInfo{
				{
					Name:        "message",
					Type:        "string",
					Description: "Message to echo back",
					Required:    true,
				},
			},
		},
		{
			ID:          "add",
			Name:        "Add Numbers",
			Description: "Adds two numbers together",
			Parameters: []actions.ParameterInfo{
				{
					Name:        "a",
					Type:        "number",
					Description: "First number",
					Required:    true,
				},
				{
					Name:        "b",
					Type:        "number",
					Description: "Second number",
					Required:    true,
				},
			},
		},
	}
}

func (e *SimpleExecutor) Execute(actionID string, params map[string]interface{}) (*types.ActionResult, error) {
	switch actionID {
	case "hello":
		name, ok := params["name"].(string)
		if !ok {
			name = "World"
		}
		return &types.ActionResult{
			Success: true,
			Message: fmt.Sprintf("Hello, %s!", name),
			Data:    map[string]interface{}{"greeting": name},
		}, nil

	case "echo":
		message, ok := params["message"].(string)
		if !ok {
			return &types.ActionResult{
				Success: false,
				Message: "Message parameter is required and must be a string",
			}, nil
		}
		return &types.ActionResult{
			Success: true,
			Message: fmt.Sprintf("Echo: %s", message),
			Data:    map[string]interface{}{"original": message, "echo": message},
		}, nil

	case "add":
		a, aOk := params["a"].(float64)
		b, bOk := params["b"].(float64)
		if !aOk || !bOk {
			return &types.ActionResult{
				Success: false,
				Message: "Both 'a' and 'b' parameters are required and must be numbers",
			}, nil
		}
		result := a + b
		return &types.ActionResult{
			Success: true,
			Message: fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result),
			Data:    map[string]interface{}{"a": a, "b": b, "result": result},
		}, nil

	default:
		return &types.ActionResult{
			Success: false,
			Message: fmt.Sprintf("Unknown action: %s", actionID),
		}, nil
	}
}

func (e *SimpleExecutor) Validate(actionID string, params map[string]interface{}) error {
	switch actionID {
	case "hello":
		// name is optional, no validation needed
		return nil

	case "echo":
		if _, ok := params["message"]; !ok {
			return fmt.Errorf("message parameter is required")
		}
		if _, ok := params["message"].(string); !ok {
			return fmt.Errorf("message parameter must be a string")
		}
		return nil

	case "add":
		if _, ok := params["a"]; !ok {
			return fmt.Errorf("parameter 'a' is required")
		}
		if _, ok := params["b"]; !ok {
			return fmt.Errorf("parameter 'b' is required")
		}
		if _, ok := params["a"].(float64); !ok {
			return fmt.Errorf("parameter 'a' must be a number")
		}
		if _, ok := params["b"].(float64); !ok {
			return fmt.Errorf("parameter 'b' must be a number")
		}
		return nil

	default:
		return fmt.Errorf("unknown action: %s", actionID)
	}
}

// Required for plugin compilation
func main() {}
