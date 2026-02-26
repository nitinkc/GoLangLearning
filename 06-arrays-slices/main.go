// Lesson 6: Arrays and Slices in Go
// Arrays are fixed-size, Slices are dynamic

package main

import "fmt"

func main() {
	// ARRAYS - Fixed size
	fmt.Println("=== Arrays ===")

	// Declare array with size
	var numbers [5]int
	fmt.Println("Empty array:", numbers) // Zero values

	// Initialize array
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("After assignment:", numbers)

	// Array literal
	fruits := [3]string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Length:", len(fruits))

	// Let compiler count
	colors := [...]string{"red", "green", "blue", "yellow"}
	fmt.Println("Colors:", colors)
	fmt.Println("Length:", len(colors))

	// SLICES - Dynamic size
	fmt.Println("\n=== Slices ===")

	// Create slice from array
	allColors := colors[:]    // All elements
	someColors := colors[1:3] // Elements 1 and 2
	fmt.Println("All colors:", allColors)
	fmt.Println("Some colors (1:3):", someColors)

	// Slice literal
	primes := []int{2, 3, 5, 7, 11}
	fmt.Println("Primes:", primes)

	// Make slice with length and capacity
	slice := make([]int, 3, 5) // length=3, capacity=5
	fmt.Printf("Slice: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// SLICE OPERATIONS
	fmt.Println("\n=== Slice Operations ===")

	// Append
	slice = append(slice, 10)
	slice = append(slice, 20, 30) // Multiple values
	fmt.Println("After append:", slice)

	// Append slice to slice
	more := []int{40, 50}
	slice = append(slice, more...)
	fmt.Println("After appending slice:", slice)

	// Copy
	source := []int{1, 2, 3}
	dest := make([]int, len(source))
	copy(dest, source)
	fmt.Println("Copied slice:", dest)

	// ITERATING
	fmt.Println("\n=== Iteration ===")

	// For loop with index
	fmt.Print("By index: ")
	for i := 0; i < len(primes); i++ {
		fmt.Print(primes[i], " ")
	}
	fmt.Println()

	// For-range
	fmt.Print("With range: ")
	for _, value := range primes {
		fmt.Print(value, " ")
	}
	fmt.Println()

	// Range with index
	fmt.Println("Index and value:")
	for i, v := range primes[:3] {
		fmt.Printf("  [%d] = %d\n", i, v)
	}

	// SLICE TRICKS
	fmt.Println("\n=== Slice Tricks ===")

	data := []int{1, 2, 3, 4, 5}

	// Remove element at index 2
	removed := append(data[:2], data[3:]...)
	fmt.Println("Remove index 2:", removed)

	// Insert element
	data = []int{1, 2, 4, 5}
	data = append(data[:2], append([]int{3}, data[2:]...)...)
	fmt.Println("Insert 3 at index 2:", data)

	// 2D SLICES
	fmt.Println("\n=== 2D Slices ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(" ", row)
	}
	fmt.Println("Element [1][2]:", matrix[1][2])
}

/*
EXERCISES:
1. Create an array of your 5 favorite movies
2. Create a slice and practice append operations
3. Write a function that reverses a slice
4. Create a 2D slice for a tic-tac-toe board

RUN THIS FILE:
  go run main.go
*/
