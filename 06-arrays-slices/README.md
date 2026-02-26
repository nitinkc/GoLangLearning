# Lesson 6: Arrays and Slices

## What You'll Learn
- Fixed-size arrays
- Dynamic slices
- Slice operations (append, copy)
- Iteration with range
- 2D slices

## Key Concepts

### Arrays (Fixed Size)
```go
var arr [5]int              // Zero initialized
arr := [3]string{"a", "b", "c"}
arr := [...]int{1, 2, 3}    // Compiler counts
```

### Slices (Dynamic)
```go
slice := []int{1, 2, 3}     // Slice literal
slice := make([]int, 5)     // Length 5
slice := make([]int, 3, 10) // Length 3, capacity 10
```

### Slice Operations
```go
slice = append(slice, 4)    // Add element
slice = append(s1, s2...)   // Append slice
copy(dest, src)             // Copy slices
```

### Slicing
```go
arr[1:3]  // Elements 1, 2
arr[:3]   // First 3 elements
arr[2:]   // From index 2 to end
arr[:]    // All elements
```

## Running the Code
```bash
go run main.go
```
