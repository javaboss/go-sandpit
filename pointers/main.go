package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a // re-use same address (pointer to int a  and assign to b)
	// short declaration would be b = &a
	fmt.Println(b)
	fmt.Println(&b) // provides the address of the pointer
	fmt.Println(*b) // dereferences the pointer and provies the value stored at that address

	*b = 43
	fmt.Println(*b) // changes the value at the pointer location

}
