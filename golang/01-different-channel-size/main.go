package main

import (
	"fmt"
)

func example1() {

	// Initialize unbuffered channel
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- i
			fmt.Printf("t: %d ", i)
		}(i)
	}

	fmt.Println()
	for i := 0; i < 100; i++ {
		a := <-c
		fmt.Printf("m: %d ", a)
	}
}

func example2() {

	// Initialize channel of size 100
	c := make(chan int, 100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- i
			fmt.Printf("t: %d ", i)
		}(i)
	}

	fmt.Println()
	for i := 0; i < 100; i++ {
		a := <-c
		fmt.Printf("m: %d ", a)
	}
}

func main() {
	example1()
	//example2()
}
