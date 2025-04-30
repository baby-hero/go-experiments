package main

import (
	"fmt"
	"reflect"
)

func TestArray() {
	fmt.Println("Test Array: It's quite same other languages")

	var arr [3]int = [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Printf("index: %d, Value: %d\n", i, v)
	}

	x := [...]int{1, 2, 3, 4, 5} // ... means length is unknown
	// access by index
	index := 2
	fmt.Println("Index 2:", x[index], len(x))
	index += 3
	fmt.Println("Index from 1 to 4:", x[1:4], len(x[1:4]), reflect.TypeOf(x[1:4]))

	// error index out of range
	// fmt.Println("Index -1", x[-1])
	// x[index] = 6
	// fmt.Println("Index 5:", x[index], len(x))

	for i, v := range x {
		fmt.Printf("index: %d, Value: %d\n", i, v)
	}
	fmt.Println("Test Array end")
}
