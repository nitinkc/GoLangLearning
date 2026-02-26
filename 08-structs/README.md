# Lesson 8: Structs

## What You'll Learn
- Defining structs
- Creating struct instances
- Methods on structs
- Embedded structs (composition)
- Pointers vs value receivers

## Key Concepts

### Defining Structs
```go
type Person struct {
    Name string
    Age  int
}
```

### Creating Instances
```go
p := Person{Name: "Alice", Age: 25}
p := Person{"Alice", 25}  // Positional
p := &Person{Name: "Bob"} // Pointer
```

### Methods
```go
// Value receiver
func (p Person) Greet() string {
    return "Hello, " + p.Name
}

// Pointer receiver (can modify)
func (p *Person) Birthday() {
    p.Age++
}
```

### Embedding
```go
type Employee struct {
    Person        // Embedded
    Salary float64
}
// emp.Name works (promoted)
```

## Running the Code
```bash
go run main.go
```
