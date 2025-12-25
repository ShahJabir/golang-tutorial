package main

import "fmt"

func changeNumber(n int) {
	n = 2
	fmt.Println("Inside changeNumber function:", n)
}

func changeNumberByPointer(n *int) {
	*n = 2
	fmt.Println("Inside changeNumberByPointer function:", *n)
}

func main() {
	num := 1
	fmt.Println("Before changing number (Default):", num)
	changeNumber(num)
	fmt.Println("After changing (changeNumber Function) number:", num)
	changeNumberByPointer(&num)
	fmt.Println("After changing (changeNumberByPointer Function) number:", num)
}
