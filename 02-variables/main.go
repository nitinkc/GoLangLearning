// Lesson 2: Variables in Go
// Learn different ways to declare and use variables

package main

import "fmt"

func main() {
	// Method 1: var keyword with type
	var name string = "Gopher"
	var age int = 25

	// Method 2: var with type inference
	var city = "San Francisco" // Go infers the type as string

	// Method 3: Short declaration (most common) - only inside functions
	country := "USA"
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
}

/*
EXERCISES:
1. Create variables for your personal info (name, age, city)
2. Try using the short declaration syntax (:=)
3. Create a constant for your favorite number
4. Observe what happens when you try to reassign a constant

RUN THIS FILE:
  go run main.go
*/
