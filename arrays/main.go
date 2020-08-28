package main

import "fmt"

func main() {

	// ARRAYS
	a := [5]int{1, 2, 3, 4, 5}

	for i, v := range a {
		fmt.Printf("%d\t%v\n", i, v)
	}

	fmt.Printf("%T\n", a)

	// NOW THE SAME, BUT WITH A SLICE

	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		fmt.Printf("%d\t%v\n", i, v)
	}

	fmt.Printf("%T\n", s)

}
