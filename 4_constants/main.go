package main

import "fmt"

const (
	PI = 3.14159
	NAME = "Shah Jabir"
	AGE = 22
)

func main() {
	const name = "John Doe"
	const age = 30
	// This will throw an error because you cannot change the value of a constant
	// name = "Jane Doe"
	fmt.Println("In scope Name:", name)
	fmt.Println("In scope Age:", age)
	fmt.Println("My name is:", NAME, "and I am", AGE, "years old.")
	fmt.Println("The value of PI is:", PI)
}
