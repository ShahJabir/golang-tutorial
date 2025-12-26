package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		println("Processed number:", num)
	}
}

func Add(result chan int, a int, b int) {
	sum := a + b
	result <- sum
}

func main() {
	numChannels := make(chan int)
	resChannels := make(chan int)

	//  Receive input to channels from Add function
	go Add(resChannels, 5, 7)
	before_resChannels := time.Now()
	sum := <-resChannels
	fmt.Println("Sum is:", sum)
	fmt.Println("Time taken for addition:", time.Since(before_resChannels).Nanoseconds(), "ns")
	//  Close resChannels after use
	close(resChannels)

	//  Sending random numbers to numChannels
	before_numChannels := time.Now()
	go processNum(numChannels)
	for i := 0; i < 100; i++ {
		numChannels <- rand.Intn(1000)
	}
	println("Time taken for random loop:", time.Since(before_numChannels).Nanoseconds(), "ns")
	//  Close numChannels after sending all numbers
	close(numChannels)
}
