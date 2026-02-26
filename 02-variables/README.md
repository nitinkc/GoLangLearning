# Lesson 2: Variables

## What You'll Learn
- Variable declaration methods
- Type inference
- Constants
- Zero values

## Key Concepts

### Declaration Methods
```go
// Explicit type
var name string = "Go"

// Type inference
var age = 25

// Short declaration (inside functions only)
count := 10
```

### Constants
```go
const pi = 3.14159
const maxSize = 100
```

### Zero Values
Variables declared without initialization get zero values:
- `int`: 0
- `string`: ""
- `bool`: false
- `float64`: 0.0

## Running the Code
```bash
go run main.go
```
