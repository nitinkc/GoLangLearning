# Lesson 7: Maps

## What You'll Learn
- Creating maps
- Accessing and modifying values
- Checking key existence
- Iterating over maps
- Nested maps

## Key Concepts

### Creating Maps
```go
// Using make
m := make(map[string]int)

// Map literal
m := map[string]int{
    "one": 1,
    "two": 2,
}
```

### Operations
```go
m["key"] = value        // Set
value := m["key"]       // Get
delete(m, "key")        // Delete
len(m)                  // Length
```

### Check Key Exists
```go
value, ok := m["key"]
if ok {
    // key exists
}
```

### Iteration
```go
for key, value := range m {
    // process key, value
}
```

## Running the Code
```bash
go run main.go
```
