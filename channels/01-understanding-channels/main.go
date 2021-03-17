package main

import "fmt"

func main() {
	/*
		Channels BLOCK! - unless you specify a buffer i.e:
		c := make(chan int, 1) // specifiy the buffer size as a 2nd parmater - this will stop the write to the channel from blocking
	*/

	c := make(chan int) // make the channel

	/*
		Channels BLOCK!
		 If you write something to a channel i.e. using c <- 42 it will block until the value being passed is taken off by another goroutine
		 The same happens with reading from a channel i.e. <- c will block until the value is read off the channel by another goroutine
	*/
	go func() {
		c <- 42 // push onto the channel
	}()

	fmt.Println(<-c)
}
