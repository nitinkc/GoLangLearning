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
}
