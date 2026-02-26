# Lesson 3: Data Types

## What You'll Learn
- Integer types (signed and unsigned)
- Floating point numbers
- Strings and runes
- Booleans
- Type conversion

## Key Concepts

### Integer Types
| Type | Size | Range |
|------|------|-------|
| int8 | 8 bits | -128 to 127 |
| int16 | 16 bits | -32768 to 32767 |
| int32 | 32 bits | -2B to 2B |
| int64 | 64 bits | Very large |
| uint8 | 8 bits | 0 to 255 |

### Floating Point
- `float32` - 32-bit precision
- `float64` - 64-bit precision (default)

### Strings
```go
s := "Hello"           // Regular string
m := `Multi
line`                  // Raw string literal
```

### Type Conversion
```go
i := 42
f := float64(i)  // Explicit conversion required
```

## Running the Code
```bash
go run main.go
```
