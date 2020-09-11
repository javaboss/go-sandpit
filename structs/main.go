package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Neil",
		last:  "Liddicott",
		age:   46,
	}

	p2 := person{
		first: "Helen",
		last:  "Liddicott",
		age:   45,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
