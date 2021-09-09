package main

import (
	"fmt"
	"runtime"
)

// EXERCISE 7 - additonal example at bottom....

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
	}

	return c
}

/* COULD DO THIS ALSO ...
c := make(chan int)

for j := 0; j < 10; j++ {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
}

for k := 0; k < 100; k++ {
	fmt.Println(k, <-c)
}

fmt.Println("about to exit")
*/
