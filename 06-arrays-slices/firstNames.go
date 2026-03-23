package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declare an array of 3 strings
	var fullNames [3]string

	// Assign values to the array
	fullNames[0] = "Alice Smith"
	fullNames[1] = "Bob Johnson"
	fullNames[2] = "Charlie Brown"

	// Print the full names
	for i, name := range fullNames {
		fmt.Println("Full Name", i+1, ":", name)
	}

	//Slice to hold first names
	firstNames := []string{}
	for _, name := range fullNames {
		firstName := strings.Fields(name)[0]
		// Append the first name to the slice
		firstNames = append(firstNames, firstName)
	}

	fmt.Println(firstNames)
}
