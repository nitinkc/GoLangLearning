# Lesson 10: Interfaces

## What You'll Learn
- Defining interfaces
- Implicit implementation
- Interface usage
- Type assertions and type switches
- Empty interface (any)

## Key Concepts

### Defining Interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Implementing (Implicit)
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
// Rectangle now implements Shape
```

### Type Assertion
```go
var s Shape = Rectangle{}
r, ok := s.(Rectangle)
if ok {
    // r is Rectangle
}
```

### Type Switch
```go
switch v := s.(type) {
case Rectangle:
    // v is Rectangle
case Circle:
    // v is Circle
}
```

### Empty Interface
```go
var anything interface{}
anything = 42
anything = "hello"
// Or use 'any' (Go 1.18+)
```

## Running the Code
```bash
go run main.go
```
