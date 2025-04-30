package main

import "fmt"

func TestPointer() {
	var x int = 1
	var y int
	var ip *int = &x
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
