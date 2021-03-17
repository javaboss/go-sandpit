package main

import "fmt"

func main() {
	c := make(chan int, 2)

	cr := make(<-chan int, 2) // receieve only channel
	cs := make(chan<- int, 2) // send only channel

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}
