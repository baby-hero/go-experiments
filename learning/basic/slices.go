package main

import "fmt"

func TestSlice() {
	fmt.Println("Test Slice")
	s := []int{1, 2, 3} // Variable size
	fmt.Println(s)

	s = append(s, 4)
	fmt.Println()
	// len and capacity
	fmt.Println("Len and capacity", s, len(s), cap(s))
	s = append(s, 5, 6, 7)
	fmt.Println("Len and capacity", s, len(s), cap(s))

	// make: create a slice (and array)/ map or channel
	s = make([]int, 3, 5) // len = 3, cap = 5
	fmt.Println(s)
	s = append(s, 9)
	fmt.Println("Len and capacity", s, len(s), cap(s))
	x := make([]float32, 10)
	fmt.Println("Len and capacity", x, len(x), cap(x))
	x = append(x, 0.1, 0.2)
	fmt.Println("Len and capacity", x, len(x), cap(x))
	y := make([]string, 0)
	fmt.Println("Len and capacity", y, len(y), cap(y))
	y = append(y, "test")
	fmt.Println("Len and capacity", y, len(y), cap(y))
	fmt.Println("Test Slice end")
}
