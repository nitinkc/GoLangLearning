package main

import "fmt"

func main() {
	// Iteration examples for slices
	fmt.Println("=== Slice Iteration Examples ===")

	primes := []int{2, 3, 5, 7, 11}

	// For loop with index
	fmt.Print("By index: ")
	for i := 0; i < len(primes); i++ {
		fmt.Print(primes[i], " ")
	}
	fmt.Println()

	// For-range
	fmt.Println("With range: ")
	for index, value := range primes {
		fmt.Printf("Index : %v :: value : %v \n", index, value)
	}
	fmt.Println()

	// Range with index
	fmt.Println("Index and value:")
	for i, v := range primes[:3] {
		fmt.Println("  [%d] = %d", i, v)
	}

	// Additional iteration patterns
	fmt.Println("\n=== Additional Iteration Patterns ===")

	// Using range to modify a slice (careful: value is a copy)
	nums := []int{10, 20, 30}
	fmt.Println("Before modify attempt:", nums)
	for _, v := range nums {
		v = v + 1 // does not modify original
	}
	fmt.Println("After modify attempt (unchanged):", nums)

	// Proper way to modify in-place using index
	for i := range nums {
		nums[i] = nums[i] + 1
	}
	fmt.Println("After in-place modify:", nums)

	// Iterating over a slice of structs
	type Point struct{ X, Y int }
	points := []Point{{1, 2}, {3, 4}}
	for i, p := range points {
		fmt.Printf("Point %d: %#v\n", i, p)
	}

	// Safe iteration while deleting: iterate and build a new slice
	vals := []int{1, 2, 3, 4, 5}
	keep := vals[:0] // reuse backing array
	for _, v := range vals {
		if v%2 == 1 { // keep odd
			keep = append(keep, v)
		}
	}
	fmt.Println("Kept odds:", keep)
}
