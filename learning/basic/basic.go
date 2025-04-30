package main

import (
	"fmt"
	"unicode/utf8"
)

type Animal struct {
	Species string
}

func (*Animal) addFloat(x float32, y float32) float32 {
	return x + y
}

func testPointer() {
	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Println(x, y, ip)
	x = 2
	fmt.Println(x, y, ip)

	ptr := new(int) // pointer to int
	fmt.Println(ptr, *ptr)
	*ptr = 1
	fmt.Println(ptr, *ptr)

	ptrStr := new(string)
	fmt.Println(ptrStr, "("+*ptrStr+")")
	*ptrStr = "test"
	fmt.Println(ptrStr, *ptrStr)
}

func testString(s string) {
	fmt.Println("testString", s)
	fmt.Println(len(s))                    // prints the number of bytes (a byte = 8 bit ~ a character)
	fmt.Println(utf8.RuneCountInString(s)) // prints the number of runes (characters)
	// In Go, a rune is an alias for int32 and represents a Unicode code point.
	// It's used to handle characters beyond basic ASCII, especially in UTF-8 encoded strings.
}

func main() {
	// 1. Variables and Constants
	var name string = "testing" //     default string on go is utf-8
	age := 30                   // Short declaration (only inside functions)
	const PI = 3.14             // Constant (unchangeable)
	fmt.Println(name, age, PI)
	age = 20
	fmt.Println(name, age, PI)

	// 2. Data Types
	var arr = [3]int{1, 2, 3} // Array
	// var slice = []int{4, 5, 6}                        // Slice
	// var dict = map[string]int{"Alice": 25, "Bob": 30} // Map
	// var num int64 = 12

	// 3. Functions
	a := Animal{
		Species: "Cat"}
	result := a.addFloat(5.3, 6.4)
	fmt.Println(result)

	// 4. If-Else
	if len(arr) > 3 {
		fmt.Printf("Arr has %d elements", len(arr))
	} else {
		fmt.Println("Testing")
	}

	// 5. Loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 6. Pointer
	testPointer()

	// 7. String
	// In Go, the default string encoding is UTF-8 (8 bits per character).
	s := "こんにちは" // Japanese for "Hello" (unicode)
	testString(s)
	s = "Hello" // English (ASCII)
	testString(s)

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
