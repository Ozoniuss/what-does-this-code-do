package main

//https://go.dev/blog/defer-panic-and-recover

import (
	"fmt"
	"sync"
)

// f just panics
func f() error {
	panic("in f")
}

func case1() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Panic in different goroutine, try to recover the panic in the goroutine
	// it occurs.
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(fmt.Errorf("panic: %v", r))

				// Signal end of execution in deferred function, if necessary.
				wg.Done()
			}
		}()

		err := f()
		if err != nil {
			fmt.Printf("panic in case 1: %v\n", err)
		}
		wg.Done()
	}()

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("case 1 done")
}

func case2() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Errorf("panic: %v", r))

			// Signal end of execution in deferred function, if necessary.
			wg.Done()
		}
	}()

	// Panic in different goroutine, try to catch the panic in the parent.
	go func() {
		err := f()
		if err != nil {
			fmt.Printf("panic in case 1: %v\n", err)
		}
		wg.Done()
	}()

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("case 2 done")
}

func main() {

	// case1()
	case2()
}
