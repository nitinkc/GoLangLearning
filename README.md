# Go Learning Project 🚀

A comprehensive, incremental Go programming course organized in easy-to-follow lessons.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.18+)
- VS Code with [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Project Structure

```
GoLangLearning/
├── 01-basics/         # Hello World, packages, imports
├── 02-variables/      # Variable declaration, types, constants
├── 03-datatypes/      # Strings, numbers, booleans, runes
├── 04-controlflow/    # If/else, switch, loops
├── 05-functions/      # Functions, parameters, returns
├── 06-arrays-slices/  # Arrays, slices, operations
├── 07-maps/           # Key-value data structures
├── 08-structs/        # Custom types, methods
├── 09-pointers/       # Memory references
├── 10-interfaces/     # Abstraction, polymorphism
├── 11-errors/         # Error handling patterns
└── 12-concurrency/    # Goroutines, channels
```

## Programming conventions and style

This project follows idiomatic Go conventions so you can read, write, and review Go code comfortably.
If you come from Java, the main differences to keep in mind are: packages and module paths are lowercase; exported (public) identifiers start with a capital letter; unexported (private) identifiers start with lowercase. Below is a compact guide and examples.

- Package names
  - Short, all-lowercase, usually the last element of the import path. Example: `package mathutil`.
  - Files belonging to the same package share the package name at the top of the file.

- Module names and `go.mod`
  - The module path (in `module` line of `go.mod`) is typically the repository import path (for example, `github.com/you/project`). Use lowercase and avoid introducing leading short prefixes like `m` unless it is part of the real path.

- Exported vs unexported identifiers
  - Any name that starts with a capital letter is exported (visible outside the package).
  - Names that start with a lowercase letter are unexported (package-private).
  - Example:

    ```go
    package greetings

    // Exported function (available to other packages)
    func PrintHello(name string) string {
        return "Hello " + name
    }

    // unexported helper (only visible inside package greetings)
    func formatName(name string) string {
        return strings.TrimSpace(name)
    }

    // Exported struct and mixed exported/unexported fields
    type Person struct {
        Name string // exported
        age  int    // unexported
    }
    ```

- Types and naming
  - Exported types and functions use PascalCase (e.g., `Person`, `NewClient`).
  - Local variables use short, descriptive names (e.g., `i`, `cfg`, `req`).
  - Use consistent casing for acronyms: `HTTPClient`, not `HttpClient`.

- File names
  - Use all-lowercase, words separated by underscores or no separator: `main.go`, `http_server.go`. Avoid CamelCase file names.

- Constants and variables
  - Prefer CamelCase for exported constants (e.g., `DefaultTimeout`) and short lowercase for local variables. Avoid ALL_CAPS or underscores like `MAX_SIZE` unless interacting with an external API that requires it.

- Interfaces
  - Name interfaces by behavior when appropriate, often with an `-er` suffix: `Reader`, `Writer`, `Formatter`.

- Comments and documentation
  - Exported functions, types, and packages should have doc comments starting with the identifier name (Godoc style):
    `// PrintHello prints a friendly greeting.`

- Idiomatic patterns
  - Error values are returned as the last return value: `func Do() (Result, error)`.
  - Short receiver names for methods: `func (p *Person) String() string`.

- Formatting and tooling
  - Run `go fmt` (or `gofmt`) to format code automatically.
  - Use `go vet` and linters (e.g., `golangci-lint`) for better code quality.

Quick checklist when writing Go code in this repo:
- [ ] package names are lowercase and short
- [ ] exported names start with a capital letter; unexported start lowercase
- [ ] file names are lowercase
- [ ] use `go fmt` before committing
- [ ] write doc comments for exported items

References and additional reading:
- Effective Go: https://golang.org/doc/effective_go
- Go Code Review Comments: https://github.com/golang/go/wiki/CodeReviewComments
- Go blog and docs: https://golang.org/doc/

## How to Use

### Running Examples

Navigate to any lesson folder and run the example:

```bash
# Example: Run the basics lesson
cd 01-basics
go run main.go
```

### Learning Path

1. **Start with lesson 01** and work through sequentially
2. **Read the README.md** in each folder for concepts
3. **Run the main.go** to see the code in action
4. **Complete the exercises** at the bottom of each file
5. **Experiment!** Modify the code and see what happens

## Quick Reference

| Command              | Description           |
|:---------------------|:----------------------|
| `go run filename.go` | Run a Go file         |
| `go build`           | Compile the program   |
| `go fmt`             | Format your code      |
| `go test`            | Run tests             |
| `go mod tidy`        | Clean up dependencies |

## Lessons Overview

### Beginner
- **01-basics**: Your first Go program
- **02-variables**: Storing data
- **03-datatypes**: Types of data in Go
- **04-controlflow**: Making decisions and loops

### Intermediate
- **05-functions**: Reusable code blocks
- **06-arrays-slices**: Collections of data
- **07-maps**: Key-value storage
- **08-structs**: Custom data types

### Advanced
- **09-pointers**: Memory and references
- **10-interfaces**: Abstraction
- **11-errors**: Handling failures gracefully
- **12-concurrency**: Parallel execution

## Tips for Learning

1. **Type the code yourself** - don't just copy/paste
2. **Make mistakes** - errors teach you more than success
3. **Experiment** - change values, add features
4. **Build something** - apply what you learn to a project

## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)

