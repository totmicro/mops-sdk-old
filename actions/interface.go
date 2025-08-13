package actions

import (
	"github.com/totmicro/mops-sdk/types"
)

// ActionExecutor represents an entity that can execute actions
type ActionExecutor interface {
	// GetID returns the unique identifier for this executor
	GetID() string

	// GetActions returns a list of actions this executor can handle
	GetActions() []ActionInfo

	// Execute performs the specified action with given parameters
	Execute(actionID string, params map[string]interface{}) (*types.ActionResult, error)

	// Validate checks if the executor can handle the given action with parameters
	Validate(actionID string, params map[string]interface{}) error
}

// ActionInfo describes an action that an executor can perform
type ActionInfo struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  []ParameterInfo        `json:"parameters,omitempty"`
	Examples    []ActionExample        `json:"examples,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Category    string                 `json:"category,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ParameterInfo describes a parameter for an action
type ParameterInfo struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default,omitempty"`
	Options     []string    `json:"options,omitempty"`
	Validation  string      `json:"validation,omitempty"`
}

// ActionExample provides usage examples for actions
type ActionExample struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
	Expected    string                 `json:"expected,omitempty"`
}

// DynamicProvider provides dynamic menu entries
type DynamicProvider interface {
	// GetID returns the unique identifier for this provider
	GetID() string

	// GetEntries generates menu entries based on the given parameter
	GetEntries(param string) ([]types.MenuEntry, error)

	// GetDescription returns a description of what this provider does
	GetDescription() string

	// SupportsParam checks if this provider can handle the given parameter
	SupportsParam(param string) bool
}

// InteractiveFunction represents a function that can be called from menus
type InteractiveFunction interface {
	// GetID returns the unique identifier for this function
	GetID() string

	// Execute runs the function with given parameters and returns a result
	Execute(params map[string]interface{}) (*types.ActionResult, error)

	// GetDescription returns a description of what this function does
	GetDescription() string

	// GetParameters returns information about the parameters this function accepts
	GetParameters() []ParameterInfo

	// Validate checks if the function can be executed with the given parameters
	Validate(params map[string]interface{}) error
}

// ActionRegistry manages action executors, providers, and functions
type ActionRegistry interface {
	// RegisterExecutor adds an action executor to the registry
	RegisterExecutor(executor ActionExecutor) error

	// RegisterProvider adds a dynamic provider to the registry
	RegisterProvider(provider DynamicProvider) error

	// RegisterFunction adds an interactive function to the registry
	RegisterFunction(function InteractiveFunction) error

	// GetExecutor retrieves an executor by ID
	GetExecutor(id string) (ActionExecutor, bool)

	// GetProvider retrieves a provider by ID
	GetProvider(id string) (DynamicProvider, bool)

	// GetFunction retrieves a function by ID
	GetFunction(id string) (InteractiveFunction, bool)

	// GetAllExecutors returns all registered executors
	GetAllExecutors() map[string]ActionExecutor

	// GetAllProviders returns all registered providers
	GetAllProviders() map[string]DynamicProvider

	// GetAllFunctions returns all registered functions
	GetAllFunctions() map[string]InteractiveFunction

	// ExecuteAction executes an action using the appropriate executor
	ExecuteAction(executorID, actionID string, params map[string]interface{}) (*types.ActionResult, error)
}
