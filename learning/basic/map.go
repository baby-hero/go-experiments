package main

import "fmt"

func TestMap() {
	fmt.Println("Test Map")
	// map is implementation of a hash table
	var idMap map[string]int
	// create map
	// idMap = make(map[string]int) // create empty map by make
	idMap = map[string]int{
		"Bob":  2,
		"John": 3,
	}

	// access map
	fmt.Println("Bob's ID:", idMap["Bob"])
	fmt.Println("John's ID:", idMap["John"])
	fmt.Println("Unknown's ID:", idMap["Unknown"]) // returns zero value of the type

	// add key
	idMap["Alice"] = 1

	// delete key
	delete(idMap, "John")

	JohnID, ok := idMap["John"]
	if !ok {
		fmt.Println("John doesn't exist")
	} else {
		fmt.Println("John's ID:", JohnID)
	}

	// loop over map
	fmt.Println("Loop over map")
	for k, v := range idMap {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Println("Test Map end")
}
