package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := append(x[:5])
	b := append(x[5:])
	c := append(x[2:7])
	d := append(x[1:6])

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
