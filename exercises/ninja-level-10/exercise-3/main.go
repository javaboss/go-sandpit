package main

import (
	"fmt"
)

func main() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // CHANELLS BLOCK!!!! Must close the channel as otherwise it will wait for something to read from it
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c { // this will block unless the channel is closed (above!)
		fmt.Println(v)
	}
}
