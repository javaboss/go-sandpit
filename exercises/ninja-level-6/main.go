package main

import "fmt"

func foo() int {
	return 10
}

func bar() (int, string) {
	return 20, "Hello"
}

func main() {
	// EXERCISE 1
	// create a func foo() that returns an int
	// create a func bar() that returns an int and a string
	f := foo()
	b1, b2 := bar()
	fmt.Println(f)
	fmt.Println(b1, b2)
}
