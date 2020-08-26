package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := []int{6, 7, 8, 9, 10}
	x = append(x, y...)
	fmt.Println(x)

	// example of deleting from a slice - take 0..2 then 7 onwards...
	x = append(x[:2], x[7:]...)
	fmt.Println(x)

	// make -- type, length, capacity
	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}
