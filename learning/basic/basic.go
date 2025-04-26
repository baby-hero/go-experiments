package main

import "fmt"

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

func main() {
	// 1. Variables and Constants
	var name string = "testing"
	age := 30       // Short declaration (only inside functions)
	const PI = 3.14 // Constant (unchangeable)
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
}
