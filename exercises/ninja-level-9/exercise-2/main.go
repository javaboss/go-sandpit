package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("Hello!", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// EXERCISE 2 - reinforce understanding of methods sets
	person := person{
		first: "Neil",
		last:  "Liddicott",
	}

	person.speak()

	saySomething(&person)
}
