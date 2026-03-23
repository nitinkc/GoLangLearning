package main

import "fmt"

func main() {
	fmt.Println("=== Runes ===")
	var r rune = 'A' // rune is alias for int32
	var emoji rune = '😀'
	fmt.Printf("Rune 'A': %c (value: %d)\n", r, r)
	fmt.Printf("Emoji: %c (value: %d)\n", emoji, emoji)
}
