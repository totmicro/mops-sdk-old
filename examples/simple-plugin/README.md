# Simple Plugin Example

This is a basic MOPS plugin that demonstrates the core SDK functionality.

## Features

- **Hello World**: Simple greeting action with optional name parameter
- **Echo Message**: Echoes back a provided message
- **Add Numbers**: Adds two numbers together

## Building

```bash
# Build as shared library
go build -buildmode=plugin -o simple-plugin.so .

# Test the build
go mod tidy
go build .
```

## Usage

Once built, the plugin provides these actions:

### Hello World
```
Action ID: hello
Parameters:
- name (string, optional): Name to greet (default: "World")
```

### Echo Message
```
Action ID: echo  
Parameters:
- message (string, required): Message to echo back
```

### Add Numbers
```
Action ID: add
Parameters:
- a (number, required): First number
- b (number, required): Second number
```

## Local Development

For local development against an unreleased SDK version, uncomment the replace directive in `go.mod`:

```go
replace github.com/totmicro/mops-sdk => ../..
```
