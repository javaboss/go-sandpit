package main

import "fmt"

// Struct
type person struct {
	first string
	last  string
	age   int
}

// Embedded Struct
type secretAgent struct {
	person
	ltk bool
}

func main() {
	// Structs
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

	// Embedded Struct
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   21,
		},
		ltk: true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	// No need to reference person when getting values from fields i.e.
	// for name collisions you can reference the type specificall i.e. sa1.person.first
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

}
