# Changelog

All notable changes to the MOPS SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2024-01-XX

### Added
- Initial release of MOPS SDK
- Core plugin interface for MOPS plugin development
- Action execution framework with parameter validation
- Comprehensive type system for MOPS operations
- Version management and API compatibility checking
- Built-in parameter and configuration validation
- Example plugins demonstrating SDK usage
- Complete API documentation and examples

### Package Structure
- `types/`: Core data structures and result types
- `actions/`: Action execution interfaces and metadata
- `plugin/`: Plugin lifecycle and configuration interfaces
- `version/`: Version management and compatibility checking

### Interfaces
- `Plugin`: Main plugin interface with lifecycle methods
- `ActionExecutor`: Interface for executing plugin actions
- `DynamicProvider`: Interface for dynamic menu providers
- `InteractiveFunction`: Interface for interactive plugin functions

### Features
- Parameter validation with type checking
- Rich action metadata with parameter descriptions
- Plugin configuration management
- Comprehensive error handling
- Support for background processes and streaming updates

[Unreleased]: https://github.com/totmicro/mops-sdk/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/totmicro/mops-sdk/releases/tag/v1.0.0
