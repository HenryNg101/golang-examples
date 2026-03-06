package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// =========================
	// Basic String
	// =========================

	s := "Hello"
	fmt.Println(s)
	fmt.Println("Length:", len(s)) // bytes
	//s[0] = 'h' -> Error

	s = "Hello 世界"
	fmt.Println(s)
	fmt.Println("Length:", len(s)) // bytes

	// Byte-type iteration
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("\n")
	// Rune-type iteration
	for i, r := range s {
		fmt.Println(i, r, string(r))
	}

	fmt.Println("Length of string (By byte):", len(s))
	fmt.Println("Length of string (By rune):", utf8.RuneCountInString(s))
}
