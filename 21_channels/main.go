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

func emailSender(emailChan <-chan string, done chan<- bool) {
	for email := range emailChan {
		fmt.Println("Sent email:", email)
	}
	done <- true
}

func main() {
	// Create channels
	numChannels := make(chan int)
	resChannels := make(chan int)
	done := make(chan bool)
	emailChannels := make(chan string, 100) // Buffered channel with capacity 100

	//  Receive input to channels from Add function
	go Add(resChannels, 5, 7)
	before_resChannels := time.Now()
	sum := <-resChannels
	fmt.Println("Sum is:", sum)
	fmt.Println("Time taken for addition:", time.Since(before_resChannels).Nanoseconds(), "ns")
	//  Close resChannels after use
	close(resChannels)

	//  Sending emails to emailChannels
	go emailSender(emailChannels, done)
	before_emailChannels := time.Now()
	for i := 1; i <= 100; i++ {
		email := fmt.Sprintf("%d@gmail.com", i)
		emailChannels <- email
	}
	fmt.Println("All emails sent to channel")
	//  Close emailChannels after sending all emails
	close(emailChannels)
	<-done
	fmt.Println("Time taken for email loop:", time.Since(before_emailChannels).Milliseconds(), "ms")

	//  Sending random numbers to numChannels
	before_numChannels := time.Now()
	go processNum(numChannels)
	for i := 1; i <= 100; i++ {
		numChannels <- rand.Intn(1000)
	}
	fmt.Println("Time taken for random loop:", time.Since(before_numChannels).Microseconds(), "ms")
	//  Close numChannels after sending all numbers
	close(numChannels)
}
