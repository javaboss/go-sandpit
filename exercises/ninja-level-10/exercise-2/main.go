package main

import (
	"fmt"
)

func main() {
	// receive only channel
	// cr := make(<-chan int)
	// send only channel
	// cr := make(chan<- int)

	// solution
	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
