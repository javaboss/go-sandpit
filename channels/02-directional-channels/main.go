package main

import "fmt"

func main() {
	c := make(chan int, 2)

	// c := make(<-chan int, 2) // RECEIVE only channel - example
	// c := make(chan<- int, 2) // SEND only channel - example

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Printf("%T\n", c)
}
