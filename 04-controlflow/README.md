# Lesson 4: Control Flow

## What You'll Learn
- If-else statements
- Switch statements
- For loops (Go's only loop)
- Break and continue

## Key Concepts

### If-Else
```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}

// With initialization
if x := getValue(); x > 10 {
    // x is scoped to this if block
}
```

### Switch
```go
switch value {
case "a":
    // code
case "b", "c":
    // multiple values
default:
    // fallback
}
```

### For Loops
```go
// Traditional
for i := 0; i < 10; i++ { }

// While-style
for condition { }

// Infinite
for { break }

// Range
for index, value := range collection { }
```

## Running the Code
```bash
go run main.go
```
