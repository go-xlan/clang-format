[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/clang-format/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/clang-format/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/clang-format)](https://pkg.go.dev/github.com/go-xlan/clang-format)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/clang-format/main.svg)](https://coveralls.io/github/go-xlan/clang-format?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/clang-format.svg)](https://github.com/go-xlan/clang-format/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/clang-format)](https://goreportcard.com/report/github.com/go-xlan/clang-format)

# clang-format

Go package for clang-format with Protocol Buffers and C/C++ batch formatting capabilities.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Key Features

🎯 **Intelligent Proto Formatting**: Smart clang-format package with Google style defaults  
⚡ **Two Operation Modes**: Both preview (DryRun) and in-place formatting support  
🔄 **Batch Processing**: Recursive project-wide .proto file formatting  
🌍 **Configurable Styles**: Customizable formatting styles with JSON configuration  
📋 **Comprehensive Logging**: Detailed operation logs with structured output

## Installation

### Get the Package

```bash
go get github.com/go-xlan/clang-format@latest
```

### Get CLI Command

```bash
go install github.com/go-xlan/clang-format/cmd/clang-format-batch@latest
```

## Prerequisites

Setup clang-format as requirement:

```bash
# macOS
brew install clang-format

# Ubuntu/Debian
sudo apt-get install clang-format

# Verify installation
clang-format --version
```

## Quick Start

### Command Line Usage

```bash
# Format all .proto files in current project
clang-format-batch --extensions=".proto"

# Format C/C++ files
clang-format-batch --extensions=".c,.cpp,.h"

# Format multiple file types
clang-format-batch --extensions=".proto,.c,.cpp,.h"

# Use short flag
clang-format-batch -e ".proto,.cc,.hh"
```

## Library Usage

### Protocol Buffers Formatting (Primary Use Case)

```go
package main

import (
    "fmt"
    "github.com/go-xlan/clang-format/protoformat"
    "github.com/yyle88/must"
    "github.com/yyle88/osexec"
    "github.com/yyle88/rese"
)

func main() {
    execConfig := osexec.NewExecConfig()
    style := protoformat.NewStyle()
    
    // Preview .proto file formatting (DryRun)
    output := rese.V1(protoformat.DryRun(execConfig, "example.proto", style))
    fmt.Println(string(output))
    
    // Format single .proto file
	rese.V1(protoformat.Format(execConfig, "example.proto", style))
    
    // Format entire project (batch processing)
    must.Done(protoformat.FormatProject(execConfig, "./proto-project", style))
}
```

### Custom Style Configuration

```go
customStyle := &clangformat.Style{
    BasedOnStyle:                "LLVM",
    IndentWidth:                 4,
    ColumnLimit:                 80,
    AlignConsecutiveAssignments: true,
}

output := rese.V1(protoformat.DryRun(execConfig, "example.proto", customStyle))
```

### General File Formatting (C/C++ Support)

```go
import "github.com/go-xlan/clang-format/clangformat"

// Format C/C++ files
output := rese.V1(clangformat.DryRun(execConfig, "example.cpp", style))
must.Done(clangformat.Format(execConfig, "example.cpp", style))
```

## API Reference

### clangformat Package

- `NewStyle()` - Creates default Google-based style configuration
- `DryRun(config, path, style)` - Preview formatting without file modification
- `Format(config, path, style)` - Use formatting on file

### protoformat Package

- `NewStyle()` - Creates Protocol Buffers optimized style configuration
- `DryRun(config, path, style)` - Preview .proto file formatting
- `Format(config, path, style)` - Format single .proto file
- `FormatProject(config, path, style)` - Batch format all .proto files in project

### Style Configuration

```go
type Style struct {
    BasedOnStyle                string // "Google", "LLVM", "Chromium", etc.
    IndentWidth                 int    // Count of spaces for indentation
    ColumnLimit                 int    // Maximum line length (0 = no limit)
    AlignConsecutiveAssignments bool   // Align assignments at assignment signs
}
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-xlan/clang-format.svg?variant=adaptive)](https://starchart.cc/go-xlan/clang-format)