package main

import "fmt"

func main() {
	foo()
	bar("Neil")

	s := woo("Helen")
	fmt.Println(s)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	// Variadic Parameters
	// you can pass in a variable number of parameters
	// if you are passing a variadic parameter it has to be the LAST paramaeter
	// ...int in the method signature denotes variadic parameters
	// to unfurl a slice of ints you do int...
	// variadic parameters also allow *nothing* to  be passed in

	sum := sum(2, 3, 4, 5, 6, 7, 9)

	/* or...
	 	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
		sum := sum(xi...)
	*/

	fmt.Println("The total is", sum)
}

func foo() {
	fmt.Println("Hell from foo")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

// return something
func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

// multiple return values
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false

	return a, b
}

// variadic parameters - passing in a set of ints creates a slide of ints
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}

	return sum
}
