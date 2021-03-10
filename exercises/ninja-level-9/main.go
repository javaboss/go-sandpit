package main

import (
	"fmt"
	"sync"
)

/*
instead of using separate functions (below) you could use anonymous functions i.e
go fund() {
	// do something
}()
*/

var wg sync.WaitGroup

func routineOne() {
	fmt.Println("ROUTINE ONE")
	wg.Done()
}

func routineTwo() {
	fmt.Println("ROUTINE TWO")
	wg.Done()
}

func main() {
	// EXERCISE 1 - Goroutines and WaitGroup

	wg.Add(2)

	go routineOne()
	go routineTwo()

	wg.Wait()
	fmt.Println("DONE")
}
