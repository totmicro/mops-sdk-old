package plugin

import (
	"github.com/totmicro/mops-sdk/actions"
)

// Plugin represents a MOPS plugin
type Plugin interface {
	// GetName returns the unique name of the plugin
	GetName() string

	// GetVersion returns the version of the plugin
	GetVersion() string

	// GetDescription returns a human-readable description of the plugin
	GetDescription() string

	// GetAPIVersion returns the MOPS API version this plugin is compatible with
	GetAPIVersion() string

	// Initialize is called when the plugin is loaded
	// The config parameter contains the plugin's configuration
	Initialize(config *PluginConfig) error

	// GetProviders returns all dynamic providers this plugin provides
	GetProviders() []actions.DynamicProvider

	// GetExecutors returns all action executors this plugin provides
	GetExecutors() []actions.ActionExecutor

	// GetFunctions returns all interactive functions this plugin provides
	GetFunctions() []actions.InteractiveFunction

	// Cleanup is called when the plugin is being unloaded
	Cleanup() error
}

// PluginConfig contains configuration for a plugin
type PluginConfig struct {
	// Name of the plugin
	Name string `yaml:"name"`

	// Version of the plugin
	Version string `yaml:"version"`

	// Whether the plugin is enabled
	Enabled bool `yaml:"enabled"`

	// Plugin-specific configuration
	Config map[string]interface{} `yaml:"config,omitempty"`

	// Environment variables for the plugin
	Environment map[string]string `yaml:"environment,omitempty"`

	// Resource limits for the plugin
	Limits *ResourceLimits `yaml:"limits,omitempty"`

	// Permissions for the plugin
	Permissions *Permissions `yaml:"permissions,omitempty"`
}

// ResourceLimits defines resource constraints for a plugin
type ResourceLimits struct {
	// Maximum memory usage in MB
	MaxMemoryMB int `yaml:"max_memory_mb,omitempty"`

	// Maximum CPU usage as percentage
	MaxCPUPercent int `yaml:"max_cpu_percent,omitempty"`

	// Maximum execution time in seconds
	MaxExecutionSeconds int `yaml:"max_execution_seconds,omitempty"`
}

// Permissions defines what a plugin is allowed to do
type Permissions struct {
	// Allow network access
	NetworkAccess bool `yaml:"network_access,omitempty"`

	// Allow file system access
	FileSystemAccess bool `yaml:"filesystem_access,omitempty"`

	// Allow executing system commands
	SystemCommands bool `yaml:"system_commands,omitempty"`

	// Allowed file paths (if filesystem access is enabled)
	AllowedPaths []string `yaml:"allowed_paths,omitempty"`

	// Allowed commands (if system commands are enabled)
	AllowedCommands []string `yaml:"allowed_commands,omitempty"`
}

// PluginMetadata contains metadata about a plugin
type PluginMetadata struct {
	// Plugin information
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Author      string `yaml:"author,omitempty"`
	URL         string `yaml:"url,omitempty"`
	License     string `yaml:"license,omitempty"`

	// MOPS compatibility
	MOPSVersionMin string `yaml:"mops_version_min"`
	MOPSVersionMax string `yaml:"mops_version_max,omitempty"`
	APIVersion     string `yaml:"api_version"`

	// Runtime requirements
	RequiresNetwork        bool     `yaml:"requires_network,omitempty"`
	RequiresSystemCommands bool     `yaml:"requires_system_commands,omitempty"`
	SystemCommands         []string `yaml:"system_commands,omitempty"`
	Dependencies           []string `yaml:"dependencies,omitempty"`

	// Build information
	BuildInfo *BuildInfo `yaml:"build_info,omitempty"`

	// Categories and tags
	Categories []string `yaml:"categories,omitempty"`
	Tags       []string `yaml:"tags,omitempty"`
}

// BuildInfo contains information about how the plugin was built
type BuildInfo struct {
	GoVersion    string            `yaml:"go_version"`
	BuildTime    string            `yaml:"build_time"`
	GitCommit    string            `yaml:"git_commit,omitempty"`
	GitBranch    string            `yaml:"git_branch,omitempty"`
	BuildFlags   []string          `yaml:"build_flags,omitempty"`
	Environment  map[string]string `yaml:"environment,omitempty"`
	TargetOS     string            `yaml:"target_os"`
	TargetArch   string            `yaml:"target_arch"`
	PluginHash   string            `yaml:"plugin_hash,omitempty"`
	Dependencies []DependencyInfo  `yaml:"dependencies,omitempty"`
}

// DependencyInfo contains information about a dependency
type DependencyInfo struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Hash    string `yaml:"hash,omitempty"`
}

// LoaderConfig holds configuration for plugin loading
type LoaderConfig struct {
	// Plugins to load - empty means load all available
	EnabledPlugins []string `yaml:"enabled_plugins,omitempty"`

	// Plugins to explicitly disable
	DisabledPlugins []string `yaml:"disabled_plugins,omitempty"`

	// Whether to fail if a plugin fails to load
	FailOnError bool `yaml:"fail_on_error,omitempty"`

	// Whether to show verbose plugin loading messages
	VerboseLogging bool `yaml:"verbose_logging,omitempty"`

	// Plugin directories to search
	PluginDirectories []string `yaml:"plugin_directories,omitempty"`

	// Plugin repositories for remote plugins
	PluginRepositories []PluginRepository `yaml:"plugin_repositories,omitempty"`

	// Security settings
	Security *SecurityConfig `yaml:"security,omitempty"`
}

// PluginRepository defines a source for plugins
type PluginRepository struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Type     string `yaml:"type"` // "github", "local", "http"
	Enabled  bool   `yaml:"enabled"`
	Priority int    `yaml:"priority"`
}

// SecurityConfig defines security settings for plugin loading
type SecurityConfig struct {
	// Allow loading unsigned plugins
	AllowUnsigned bool `yaml:"allow_unsigned,omitempty"`

	// Require specific signatures
	RequiredSignatures []string `yaml:"required_signatures,omitempty"`

	// Sandbox plugins
	EnableSandbox bool `yaml:"enable_sandbox,omitempty"`

	// Default permissions for plugins
	DefaultPermissions *Permissions `yaml:"default_permissions,omitempty"`
}

// NewPluginFunc is the signature for the plugin constructor function
// that all plugins must export
type NewPluginFunc func() Plugin

// DefaultLoaderConfig returns a default configuration for plugin loading
func DefaultLoaderConfig() *LoaderConfig {
	return &LoaderConfig{
		EnabledPlugins:  []string{}, // Empty means all
		DisabledPlugins: []string{},
		FailOnError:     false, // Don't fail if optional plugins can't load
		VerboseLogging:  false, // Quiet by default
		PluginDirectories: []string{
			"./plugins",
			"~/.mops/plugins",
		},
		Security: &SecurityConfig{
			AllowUnsigned:  true, // For development
			EnableSandbox:  false,
			DefaultPermissions: &Permissions{
				NetworkAccess:    true,
				FileSystemAccess: true,
				SystemCommands:   true,
			},
		},
	}
}
