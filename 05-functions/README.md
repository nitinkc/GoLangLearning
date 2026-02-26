# Lesson 5: Functions

## What You'll Learn
- Function declaration
- Parameters and return values
- Multiple returns
- Variadic functions
- Functions as values
- Closures

## Key Concepts

### Basic Function
```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

### Return Values
```go
func add(a, b int) int {
    return a + b
}

// Multiple returns
func divide(a, b int) (int, error) {
    return a / b, nil
}

// Named returns
func calc() (result int) {
    result = 42
    return
}
```

### Variadic Functions
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### Closures
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Running the Code
```bash
go run main.go
```
