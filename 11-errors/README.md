# Lesson 11: Error Handling

## What You'll Learn
- Go's error handling philosophy
- Creating and returning errors
- Custom error types
- Error wrapping (Go 1.13+)
- Panic and recover

## Key Concepts

### Basic Error Handling
```go
result, err := someFunction()
if err != nil {
    // Handle error
    return err
}
// Use result
```

### Creating Errors
```go
// Simple error
err := errors.New("something went wrong")

// Formatted error
err := fmt.Errorf("failed to process %s", item)
```

### Custom Error Types
```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return e.Message
}
```

### Error Wrapping
```go
// Wrap error with context
return fmt.Errorf("operation failed: %w", err)

// Check wrapped errors
if errors.Is(err, SomeError) { }
if errors.As(err, &myErr) { }
```

### Panic and Recover
```go
defer func() {
    if r := recover(); r != nil {
        // Handle panic
    }
}()
panic("critical error")
```

## Running the Code
```bash
go run main.go
```
