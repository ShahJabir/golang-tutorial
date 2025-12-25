package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func getLanguage() (string, string, string, string, string, string) {
	return "C", "C++", "Python", "Javascript", "Go", "Rust"
}

func main() {
	fmt.Println("Enter Two Integer:")
	var a, b int
	fmt.Scanln(&a, &b)
	result := add(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
	result = subtract(a, b)
	fmt.Printf("The difference of %d and %d is %d\n", a, b, result)
	result = multiply(a, b)
	fmt.Printf("The product of %d and %d is %d\n", a, b, result)
	result = divide(a, b)
	if b != 0 {
		fmt.Printf("The quotient of %d and %d is %d\n", a, b, result)
	}
	lang1, lang2, lang3, lang4, lang5, lang6 := getLanguage()
	fmt.Println("Programming Languages:")
	fmt.Println(lang1)
	fmt.Println(lang2)
	fmt.Println(lang3)
	fmt.Println(lang4)
	fmt.Println(lang5)
	fmt.Println(lang6)
}
