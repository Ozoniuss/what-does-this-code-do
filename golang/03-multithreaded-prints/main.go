package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var strlen = 100000

func generateLongStrings(max int) (string, string) {
	out1 := ""
	out2 := ""
	for i := 0; i < max; i++ {
		out1 += "a"
		out2 += "b"
	}
	return out1, out2
}

// printMultiThreaded attempts to print large texts to the standard output from
// multiple goroutines at (approximately) the same time.
func printMultiThreaded() {

	out1, out2 := generateLongStrings(strlen)

	w := &sync.WaitGroup{}
	w.Add(1)
	print(out1)
	go func() {
		print(out2)
		w.Done()
	}()
	w.Wait()
}

// fmtMultiThreaded attempts to print large texts to the standard output from
// multiple goroutines at (approximately) the same time, using the fmt library.
func fmtMultiThreaded() {

	out1, out2 := generateLongStrings(strlen)

	w := &sync.WaitGroup{}
	w.Add(1)

	fmt.Println(out1)
	go func() {
		fmt.Println(out2)
		w.Done()
	}()

	w.Wait()
}

// fmtMultiThreaded attempts to print large texts to the standard output from
// multiple goroutines at (approximately) the same time, using different
// loggers.
func logMultiThreaded() {
	// Printing from a single logger is thread safe (see the implementation of
	// log.Output), but printing from multiple loggers is not as clear.
	l1 := log.New(os.Stdout, "[log 1] ", log.Lmsgprefix)
	l2 := log.New(os.Stdout, "[log 2] ", log.Lmsgprefix)

	out1, out2 := generateLongStrings(strlen)

	w := &sync.WaitGroup{}
	w.Add(1)

	l1.Println(out1)
	go func() {
		l2.Println(out2)
		w.Done()
	}()

	w.Wait()
}

func main() {
	printMultiThreaded()
	// fmtMultiThreaded()
	// logMultiThreaded()
}
