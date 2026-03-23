package main

import "fmt"

func main() {
	fmt.Println("=== Integers ===")
	var i int = 42        // Platform dependent (32 or 64 bit)
	var i8 int8 = 127     // -128 to 127
	var i16 int16 = 32767 // -32768 to 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	fmt.Printf("int: %d, int8: %d, int16: %d\n", i, i8, i16)
	fmt.Printf("int32: %d, int64: %d\n", i32, i64)

	// Unsigned integers (no negative values)
	var ui uint = 42
	var ui8 uint8 = 255 // 0 to 255 (also called byte)
	fmt.Printf("uint: %d, uint8/byte: %d\n", ui, ui8)
}
