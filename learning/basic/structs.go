package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Phone string
}

func TestStruct() {
	fmt.Println("Test Struct")
	// initialize struct
	q := new(Person) // init with zero value
	fmt.Println(q)

	p := Person{
		Name:  "John",
		Age:   30,
		Phone: "123-456-7890",
	}
	fmt.Println(p)
	fmt.Println("Test Struct end")
}
