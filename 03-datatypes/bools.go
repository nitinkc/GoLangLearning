package main

import "fmt"

func main() {
	fmt.Println("\n=== Booleans ===")
	var isTrue bool = true
	var isFalse bool = false
	fmt.Printf("isTrue: %t, isFalse: %t\n", isTrue, isFalse)

	// Boolean operations
	fmt.Printf("AND: %t, OR: %t, NOT: %t\n", isTrue && isFalse, isTrue || isFalse, !isTrue)
}
