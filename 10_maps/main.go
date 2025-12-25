package main

import "fmt"

func main() {
	// golang maps are similar to hash tables or dictionaries in other languages.
	// They store key-value pairs and provide fast lookups, additions, and deletions.
	m := make(map[string]int)

	// Adding key-value pairs to the map
	m["apple"] = 5
	m["banana"] = 3
	m["orange"] = 7
	fmt.Println("Map after adding items:", m)

	// Accessing values by key
	appleCount := m["apple"]
	fmt.Println("Number of apples:", appleCount)

	// Checking if a key exists
	bananaCount, exists := m["banana"]
	if exists {
		fmt.Println("Number of bananas:", bananaCount)
	} else {
		fmt.Println("Bananas not found in the map.")
	}
}
