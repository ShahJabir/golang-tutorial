package main

import "fmt"

func main() {
	fmt.Println("Array Example in Go")
	// Declare and initialize an array of integers
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array of integers:", numbers)
	// Accessing elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])
	// Modifying an element
	numbers[1] = 25
	fmt.Println("Modified array:", numbers)
	// Length of the array
	fmt.Println("Length of the array:", len(numbers))
	// Iterating over the array
	fmt.Println("Iterating over the array:")
	for i, v := range numbers {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}
	// Different types of arrays
	strings := [3]string{"Go", "is", "fun"}
	fmt.Println("Array of strings:", strings)
	floats := [4]float64{3.14, 1.618, 2.718, 0.577}
	fmt.Println("Array of floats:", floats)
	boolArray := [2]bool{true, false}
	fmt.Println("Array of booleans:", boolArray)
}
