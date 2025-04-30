package main

import (
	"fmt"
	"unicode/utf8"
)

func countStringSize(s string) {
	fmt.Println("testString", s)
	fmt.Println("String size is", len(s))                          // prints the number of bytes (a byte = 8 bit ~ a character)
	fmt.Println("String runes size is", utf8.RuneCountInString(s)) // prints the number of runes (characters)
	// In Go, a rune is an alias for int32 and represents a Unicode code point.
	// It's used to handle characters beyond basic ASCII, especially in UTF-8 encoded strings.
}

func TestString() {
	s := "こんにちは" // Japanese for "Hello" (unicode)
	countStringSize(s)
	s = "Hello" // English (ASCII)
	countStringSize(s)

	s = "café"
	fmt.Println(s, len(s)) // 5 bytes: 'c','a','f','é' (é is 2 bytes in UTF-8)

	for i, r := range s {
		fmt.Printf("%d: %c (%U)\n", i, r, r)
	}
	s = "🐍Go"
	runes := []rune(s)
	fmt.Println(s, len(runes)) // 3 (🐍, G, o)
	fmt.Println(s, len(s))
}
