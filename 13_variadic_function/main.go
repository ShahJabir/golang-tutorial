package main

import "fmt"

func add(nums ...int) int {
	if len(nums) == 0 {
		fmt.Println("No numbers provided for addition.")
		return 0
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sub(nums ...int) int {
	if len(nums) == 0 {
		fmt.Println("No numbers provided for subtraction.")
		return 0
	}
	total := 0
	for _, num := range nums {
		total -= num
	}
	return total
}

func mul(nums ...int) int {
	if len(nums) == 0 {
		fmt.Println("No numbers provided for multiplication.")
		return 0
	}
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func div(nums ...int) int {
	if len(nums) == 0 {
		fmt.Println("No numbers provided for division.")
		return 0
	}
	total := nums[0]
	for _, num := range nums[1:] {
		if num != 0 {
			total /= num
		}
	}
	return total
}

func main() {
	var n int
	fmt.Print("Enter number of values: ")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Printf("Enter %d numbers: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println("Addition:", add(nums...))
	fmt.Println("Subtraction:", sub(nums...))
	fmt.Println("Multiplication:", mul(nums...))
	fmt.Println("Division:", div(nums...))
}
