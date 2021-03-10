package main

import (
	"fmt"
	"sync"
)

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
