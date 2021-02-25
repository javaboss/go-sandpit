package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	(*p).name = "Helen Liddicott"
	p.age = 45 // this is shorthand for (*p).age
}

func main() {
	// EXERCISE 1
	temperature := 21

	fmt.Println("temp = ", temperature)
	fmt.Println("temp addr = ", &temperature)

	// EXERCISE 2
	p1 := person{
		name: "Neil Liddicott",
		age:  47,
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}
