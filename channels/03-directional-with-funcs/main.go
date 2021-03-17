package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go send(c)

	// receive
	receive(c) // could use waitgroups - however, not executing this func as a goroutine ensures that it blocks until send() is complete

	fmt.Println("about to exit ")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
