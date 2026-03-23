// Lesson 2: Variables in Go
// Learn different ways to declare and use variables

package main

import "fmt"

func main() {
	// HAS TO BE USED : This variable is declared but not used, it will cause a compile error
	//var unusedVariable int

	// Method 1: var keyword with type
	var name string = "Gopher"
	var age int = 25

	// Method 2: var with type inference, Assign IMMEDIATELY
	var city = "San Francisco" // Go infers the type as string

	// Method 3: Short declaration (most common) - only inside functions
	country := "USA" //Constants can't be declared this way, only variables
	score := 95.5

	// Multiple variable declaration
	var (
		firstName = "John"
		lastName  = "Doe"
		isActive  = true
	)

	// Print all variables
	fmt.Println("=== Variable Examples ===")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Score:", score)
	fmt.Printf("Full Name: %s %s, Active: %t\n", firstName, lastName, isActive)

	// Constants - values that cannot change
	const pi = 3.14159
	const greeting = "Hello"
	fmt.Println("\n=== Constants ===")
	fmt.Println("Pi:", pi)
	fmt.Println("Greeting:", greeting)

	// Reassigning variables
	age = 26 // OK - variables can be reassigned
	fmt.Println("\nUpdated Age:", age)

	// Zero values - Go initializes variables with default values
	var defaultInt int
	var defaultString string
	var defaultBool bool
	var defaultFloat float64

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("int: %d, string: '%s', bool: %t, float: %f\n",
		defaultInt, defaultString, defaultBool, defaultFloat)

	//Print TYPE
	fmt.Printf("int: %T, string: '%T', bool: %T, float: %T\n",
		defaultInt, defaultString, defaultBool, defaultFloat)
}

/*
RUN THIS FILE:
  go run main.go
*/
