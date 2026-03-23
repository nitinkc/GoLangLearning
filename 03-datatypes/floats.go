package main

import "fmt"

func main() {
	fmt.Println("\n=== Floating Point ===")
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	fmt.Printf("float32: %f, float64: %.15f\n", f32, f64)
}
