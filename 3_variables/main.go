package main

import "fmt"

func main() {
	var a, b int
	name := "Shah Jabir"
	age := 22
	// Golang is dynamic typed language
	// But we can also use static typing
	//  Golang is strongly typed language
	// This variable will throw an error if we try to assign a  integer to an string variable
	// name = 12
	fmt.Println("My name is :", name, "and I am", age, "years old.")
	fmt.Println("Enter two integers:")
	fmt.Scanln(&a, &b)
	fmt.Println("The sum of", a, "and", b, "is", a+b)
	fmt.Println("The difference between", a, "and", b, "is", a-b)
	fmt.Println("The product of", a, "and", b, "is", a*b)
	fmt.Println("The quotient of", a, "and", b, "is", a/b)
}
