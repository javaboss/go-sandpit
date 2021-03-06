package main

import (
	"fmt"
	"math"
)

var g func() string

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

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 { // this makes the circle implicitly  implement the shape interface
	return math.Pi * (c.radius * c.radius)
}

func (s square) area() float64 { // same for the  square
	return s.length * s.length
}

func info(s shape) {
	fmt.Println("The area is ", s.area())
}

func newFunc() func() string {
	// you could always just do: return func() string {}

	f := func() string {
		return "Here's your func!"
	}

	return f
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
	create a type SQUARE
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

	sq := square{
		length: 5,
	}

	cir := circle{
		100,
	}

	info(sq)
	info(cir)

	/* EXERCISE 6
	build and use an Anonymous function */
	func(x string) {
		fmt.Println("The string is", x)
	}("Mr String")

	/* EXERCISE 7
	Assign a func to a variable then call that func */
	funky := func() string {
		return "Mr String"
	}
	fmt.Printf("%T\n", funky)
	fmt.Println(funky())
	g = funky // functiona are 1st class citizens - so you can create a func and assign another func to it

	/* EXERCISE 8
	create a func that returns a func
	assign the returned func to a variable
	call the returned func */
	newFunky := newFunc()
	fmt.Println(newFunky())

	/* EXERCISE 9
	Uee a "Callback".  Create a func and pass it to a func as an argument */
	cally := func() string {
		return "This is a callback"
	}
	fmt.Println(myFunc(cally))

	/* CLOSURE
	create a closure that encloses the scope of a variable */
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func myFunc(f func() string) string {
	return f()
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
