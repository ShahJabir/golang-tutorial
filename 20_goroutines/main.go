package main

import (
	"fmt"
	"time"
)

func goroutineExample(i int) {
		go fmt.Println(i)
}

func main() {
	before_func := time.Now()
	for i:=0; i<=100; i++ {
		goroutineExample(i)
	}
	fmt.Println("Time taken (Go Routines):", time.Since(before_func))
}
