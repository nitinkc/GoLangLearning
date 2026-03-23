package main

import "fmt"

func main() {
	fmt.Println("\n=== Strings ===")
	var greeting string = "Hello, Go!"
	multiLine := `This is a
multi-line string
using backticks`

	fmt.Println(greeting)
	fmt.Println(multiLine)
	fmt.Println("String length:", len(greeting))

	// String concatenation
	first := "Hello"
	second := "World"
	combined := first + " " + second
	fmt.Println("Combined:", combined)
}
