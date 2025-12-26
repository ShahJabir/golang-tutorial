package main

import "fmt"

func Add[T int | float64](a, b T) T {
	return a + b
}

func Subtract[T int | float64](a, b T) T {
	return a - b
}

func Multiply[T int | float64](a, b T) T {
	return a * b
}

func Divide[T int | float64](a, b T) T {
	if b == 0 {
		panic("Division by zero")
	}
	return a / b
}

type Stack[T any] struct {
	elements []T
}

func main() {
	var a, b float64
	nums := Stack[float64]{
		elements: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
	}
	fmt.Println("Stack elements:", nums.elements)
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&a, &b)
	fmt.Printf("Addition: %v\n", Add(a, b))
	fmt.Printf("Subtraction: %v\n", Subtract(a, b))
	fmt.Printf("Multiplication: %v\n", Multiply(a, b))
	fmt.Printf("Division: %v\n", Divide(a, b))
}
