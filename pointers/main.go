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
	fmt.Println(*b)
}
