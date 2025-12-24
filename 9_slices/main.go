package main

import (
	"fmt"
	"slices"
)

func main() {
	// Creating and initializing a slice of strings
	fruits := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	fmt.Println("Fruits Slice:", fruits)

	numbers := make([]int, 2, 5)
	numbers = append(numbers, 10, 20)
	fmt.Println("Numbers Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	demo := []int{1, 2, 3, 4, 5}
	demo2 := []int{1, 2, 3, 4, 5}

	fmt.Println("Are two slices Equal:", slices.Equal(demo, demo2))

	demo2[4] = 6
	fmt.Println("Are two slices Equal after modification:", slices.Equal(demo, demo2))

}
