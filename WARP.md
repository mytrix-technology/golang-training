# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Repository Overview

This is a comprehensive Go training repository organized into structured learning modules, from basic concepts to expert-level implementations. The codebase uses Go 1.24.0 and includes practical examples, advanced projects, and performance optimization studies.

## Architecture

The repository follows a hierarchical learning structure:

- **A.Basic/**: Foundational Go concepts organized into 9 training modules
  - training1: Syntax basics (variables, data types, operators)
  - training2: Control structures (arrays, slices, maps, loops, branching)
  - training3: Object-oriented concepts (functions, methods, interfaces, pointers, reflection)
  - training4: Concurrency (goroutines, channels, contexts)
  - training5: Standard library utilities (strings, time, random)
  - training6: System programming (files, CLI, regex, hashing)
  - training7: Web development (HTTP, JSON, APIs)
  - training8: Data persistence (SQL, MongoDB, unit testing)
  - training9: Advanced patterns (concurrency pipelines, generics, mutexes)

- **B.Intermediate/**: Currently empty, planned for intermediate concepts
- **C.Advanced/**: Currently empty, planned for advanced topics
- **D.Expert/**: Advanced implementations
  - AI/ML integrations (Model Context Protocol, LangChain)
  - Blockchain (P2P networks)
  - Machine learning and data science modules

- **E.Other/**: Specialized topics and examples
  - Performance benchmarking and optimization
  - One Billion Row Challenge implementations
  - Health check patterns
  - Error handling patterns
  - Profile-guided optimization

## Common Development Tasks

### Running Individual Examples
Most files in the Basic training modules are standalone examples with `main()` functions:
```bash
go run A.Basic/training1/introduction/main.go
go run A.Basic/training3/function/basic-func/implement-func.go
```

### Running Benchmarks
Execute performance benchmarks to study optimization techniques:
```bash
go test -bench=. ./E.Other/benchmark/
go test -bench=BenchmarkAdd ./E.Other/benchmark/
```

### Running Specific Applications
For complex projects in D.Expert and E.Other:
```bash
# Run LangChain Gin application
go run D.Expert/large-language-models/gin-langchain/main.go

# Run One Billion Row Challenge implementations
go run E.Other/one_brc/brc2/main.go -revision=1 measurements.txt
```

### Testing
Run tests for specific modules:
```bash
go test ./E.Other/one_brc/brc4/
go test -v ./E.Other/one_brc/brc4/
```

### Building
The module can be tidied and dependencies managed:
```bash
go mod tidy
go mod download
```

## Key Architectural Patterns

### Learning Structure
Each training module contains focused examples that build upon previous concepts. Examples are self-contained with descriptive filenames that indicate their purpose.

### Performance Focus
The repository emphasizes performance optimization:
- Benchmark suites in `E.Other/benchmark/`
- One Billion Row Challenge implementations demonstrating various optimization techniques
- Profile-guided optimization examples

### Practical Applications
Expert-level modules demonstrate real-world integrations:
- **Model Context Protocol**: AI tool integration patterns
- **LangChain Integration**: LLM application development
- **P2P Blockchain**: Distributed systems implementation
- **Health Check Patterns**: Multiple health checking library implementations

### Error Handling
Comprehensive error handling patterns demonstrated across modules, with dedicated examples in `E.Other/error-handling/`.

## Module-Specific Notes

### Basic Training Modules
- Each subdirectory typically contains multiple `.go` files with main functions
- Files cannot be tested together as packages due to multiple `main()` functions
- Run individual files to see specific concept demonstrations

### Expert Modules
- Complete applications with proper package structure
- May require external dependencies (defined in go.mod)
- Include more complex architectural patterns

### Benchmark Modules
- Use Go's built-in benchmark framework
- Focus on performance optimization techniques
- Include memory allocation profiling

## Dependencies

The project uses numerous external libraries for different training scenarios:
- **Web frameworks**: Gin, Chi
- **AI/ML**: LangChain Go, MCP Golang
- **Testing**: Testify, SQL Mock
- **Health checks**: Multiple health check libraries for comparison
- **Databases**: SQLite, MongoDB drivers
- **System**: SSH server, Discord bot integration

## Development Environment

- Go 1.24.0 required
- No special build tools or Makefiles needed
- Standard Go toolchain sufficient for all operations
- Uses standard Go module structure with comprehensive dependency management
