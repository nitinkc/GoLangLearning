package main

import "fmt"

func main() {
	fmt.Println("=== Type Conversion ===")
	var intNum int = 42
	var floatNum float64 = float64(intNum) // Must be explicit
	var backToInt int = int(floatNum)
	fmt.Printf("int: %d -> float64: %f -> int: %d\n", intNum, floatNum, backToInt)
}
