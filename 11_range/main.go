package main

import "fmt"

func main() {
	// The range keyword in Go is used to iterate over elements in various data structures.
	// It's particularly useful for iterating over slices, maps, and strings.

	// Example 1: Iterating over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Example 2: Iterating over a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Example 3: Iterating over a string
	text := "Hello"
	for i, char := range text {
		fmt.Printf("Index: %d, Character: %c\n", i, char)
	}
}
