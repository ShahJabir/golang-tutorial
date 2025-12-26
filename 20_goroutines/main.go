package main

import (
	"fmt"
	"time"
)

func goroutineExample() {
	before_func := time.Now()
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	fmt.Println("Time taken (Go Routines):", time.Since(before_func))
}

func main() {
	goroutineExample()
}
