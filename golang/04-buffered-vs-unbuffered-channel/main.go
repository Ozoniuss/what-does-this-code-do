package main

import "fmt"

func runBuffered() {
	c := make(chan int, 1)

	c <- 10
	number := <-c

	fmt.Printf("number: %d", number)
}

func runUnbuffered() {
	c := make(chan int)

	c <- 10
	number := <-c

	fmt.Printf("number: %d", number)
}

func main() {

	runBuffered()
	// runUnbuffered()
}
