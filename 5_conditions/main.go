package main

import "fmt"

func main() {
	fmt.Println("Student Grade Calculator")
	fmt.Println("Please enter your score (0-100):")
	var score int
	fmt.Scanln(&score)
	if score >= 80 {
		fmt.Println("Grade: A+")
	} else if score >= 70 {
		fmt.Println("Grade: A")
	} else if score >= 60 {
		fmt.Println("Grade: A-")
	} else if score >= 50 {
		fmt.Println("Grade: B")
	} else if score >= 40 {
		fmt.Println("Grade: C")
	} else if score >= 33 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}
