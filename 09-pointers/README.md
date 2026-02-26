# Lesson 9: Pointers

## What You'll Learn
- Pointer basics
- Address and dereference operators
- Pointers in functions
- Pointers with structs
- When to use pointers

## Key Concepts

### Pointer Basics
```go
x := 42
p := &x       // Get address
fmt.Println(*p) // Dereference (42)
*p = 100      // Modify through pointer
```

### Zero Value
```go
var p *int    // nil
p = new(int)  // Allocates, returns pointer
```

### Functions
```go
// Pass by reference
func modify(p *int) {
    *p = 100  // Changes original
}

num := 42
modify(&num)  // num is now 100
```

### Structs
```go
person := &Person{Name: "Alice"}
person.Name = "Bob"  // Auto-dereference
```

### When to Use
- Modify function parameters
- Avoid copying large structs
- Represent "no value" (nil)

## Running the Code
```bash
go run main.go
```
