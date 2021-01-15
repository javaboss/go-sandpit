package main

import "fmt"

func foo() int {
	return 10
}

func bar() (int, string) {
	return 20, "Hello"
}

// takes a variadic parameter
func foo1(variadics ...int) int {
	return 42
}

func main() {
	// EXERCISE 1
	// create a func foo() that returns an int
	// create a func bar() that returns an int and a string
	f := foo()
	b1, b2 := bar()
	fmt.Println(f)
	fmt.Println(b1, b2)

	// Exercise 2
	// create a func with an identifier foo1 that take a variadic parameter of type int
	// pass in a value of type []int and unfurl the slice of int
	sliceOfInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // create a slice of int
	total := foo1(sliceOfInts...)                   // unfurl slice when passing to functions
	fmt.Println(total)
}
