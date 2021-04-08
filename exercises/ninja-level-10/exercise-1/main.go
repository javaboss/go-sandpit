package main

import (
	"fmt"
)

func main() {

	/*
		In this case, you need to use a go routine and execute an annonymous function to stop the channel blocking.
		Otherwise, use a buffered channel i.e. c := make (chan int, 1) - the go routine is no longer needed
	*/

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
