package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	Increament := counter()
	fmt.Println(Increament()) // Output: 1
	fmt.Println(Increament()) // Output: 2
	fmt.Println(Increament()) // Output: 3
}
