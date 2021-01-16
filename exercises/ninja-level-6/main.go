package main

import "fmt"

func foo() int {
	return 10
}

func bar() (int, string) {
	return 20, "Hello"
}

// takes a variadic parameter
// ranges over the values
func foo1(variadics ...int) int {
	total := 0
	for _, v := range variadics {
		total += v
	}
	return total
}

func bar1(sliceOfInt []int) int {
	total := 0
	for _, v := range sliceOfInt {
		total += v
	}
	return total
}

func one() {
	fmt.Println("One")
	defer two() // this defers the execution of this function until the enclosing function has completed
	fmt.Println("Three")
}

func two() {
	fmt.Println("Two")
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "my age is", p.age)
}

func main() {
	// EXERCISE 1
	// create a func foo() that returns an int
	// create a func bar() that returns an int and a string
	f := foo()
	b1, b2 := bar()
	fmt.Println(f)
	fmt.Println(b1, b2)

	// EXERCISE 2
	// create a func with an identifier foo1 that take a variadic parameter of type int
	// pass in a value of type []int and unfurl the slice of int
	sliceOfInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // create a slice of int
	total := foo1(sliceOfInts...)                   // unfurl slice when passing to functions
	fmt.Println(total)

	// do the same - but with a slice of ints... no need to unfurl the variadic parameters
	total2 := bar1(sliceOfInts)
	fmt.Println(total2)

	// EXERCISE 3
	// create functions using 'defer'
	one()

	// EXERCISE 4
	// Create a user defined struct with identifer 'person' and fields firstname, lastname & age
	// add a method to the type 'person' called 'speak' that outputs their name and age
	neil := person{
		first: "Neli",
		last:  "Liddicott",
		age:   47,
	}
	neil.speak()

	/* EXERCISE 5
	create a type CIRCLE
	  attach a method to each that calculates AREA and returns it
	  circle area= π r 2
	  square area = L * W
	  create a type SHAPE that defines an interface as anything that has the AREA method
	  create a func INFO which takes type shape and then prints the area
	  create a value of type square
	  create a value of type circle
	  use func info to print the area of square
	  use func info to print the area of circle
	*/
}
