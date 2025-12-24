package main

import "fmt"

func main() {
	fmt.Println("Student Grade Calculator")
	fmt.Println("Please enter your score (0-100):")
	var score int
	fmt.Scanln(&score)
	switch {
	case score < 0 || score > 100:
		fmt.Println("Invalid score")
	case score >= 80:
		fmt.Println("Grade: A+")
	case score >= 70:
		fmt.Println("Grade: A")
	case score >= 60:
		fmt.Println("Grade: A-")
	case score >= 50:
		fmt.Println("Grade: B")
	case score >= 40:
		fmt.Println("Grade: C")
	case score > 33:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
	whoamI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I am an integer: %d\n", v)
		case float64:
			fmt.Printf("I am a float: %f\n", v)
		case string:
			fmt.Printf("I am a string: %s\n", v)
		case bool:
			fmt.Printf("I am a boolean: %t\n", v)
		default:
			fmt.Printf("I am of a different type ${}: %T\n", v)
		}
	}
	whoamI(42)
	whoamI(3.14)
	whoamI("Hello, Go!")
	whoamI(true)
	whoamI([]int{1, 2, 3})
}
