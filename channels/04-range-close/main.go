package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go send(c)

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit ")
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // this closes the underlying channel... remember the pass by value here is a reference type to an underlying channel
}
