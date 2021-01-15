package main

import "fmt"

// func (r receiver) identifier(parameters) (returns(s)) {...}
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak ")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// INTERFACES
type human interface {
	speak()
}

func barbar(h human) {

	// You cannot just access h.first or h.last because you don't know which instance you have been passed
	// you need to use assertion
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barbar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barbar", h.(secretAgent).first)
	}
}

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

	/* or... to unfurl the slice
	 	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
		sum := sum(xi...)
	*/

	fmt.Println("The total is", sum)

	// DEFER when adding defer the function is executed when the enclosing function completes
	// defer one()
	one()
	two()

	// METHODS
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(p1)

	// INTERFACES
	barbar(sa1)
	barbar(sa2)
	barbar(p1)

	// ANONYMOUS FUNCTIONS
	func(x int) {
		fmt.Println("The Meaning of Life", x)
	}(42)

	// FUNC EXPRESSION
	f := func(x int) {
		fmt.Println("The year big brother stared watcing", x)
	}
	f(1984)

	// RETURNING A FUNCTION
	xy := returnAFunc()
	fmt.Printf("%T\n", xy)
	i := xy()
	fmt.Println(i)

	// this could also be done as...
	fmt.Println(returnAFunc()())

	// CALLBACKS
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	allSum := sumIt(allNumbers...)
	fmt.Println("All numbers", allSum)

	evenSum := evenOddSum(sumIt, false, allNumbers...)
	fmt.Println("Even Numnbers", evenSum)

	oddSum := evenOddSum(sumIt, true, allNumbers...)
	fmt.Println("Odd Numnbers", oddSum)

	// CLOSURES
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

	// RECURSION
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println("FACTORIAL=", n)
	// or use a loop
	n2 := loopFact(4)
	fmt.Println("LOOP FACTORIAL=", n2)
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

func one() {
	fmt.Println("One")
}

func two() {
	fmt.Println("Two")
}

func returnAFunc() func() int {
	return func() int {
		return 451
	}
}

// CALL BACKS - used when a function is passed into a function as an argument
func sumIt(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func evenOddSum(f func(x ...int) int, isOdd bool, y ...int) int {
	var xi []int
	for _, v := range y {
		if !isOdd {
			if v%2 == 0 {
				xi = append(xi, v)
			}
		} else {
			if v%2 != 0 {
				xi = append(xi, v)
			}
		}
	}
	total := f(xi...) // callback is made here
	return total
}

// CLOSURES
// a closure is a function a function value that references variables from outside it's body
// the function may access and assign to the referenced variables
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// RECUSION - FACTORIALS is an exmaple
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// Rather than recursion, you could use a loop, for example:
func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}

	return total
}
