package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p person) {

}

func main() {
	temperature := 21

	fmt.Println("temp = ", temperature)
	fmt.Println("temp addr = ", &temperature)
}
