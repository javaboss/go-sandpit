package main

import "fmt"

type person struct {
	first      string
	last       string
	favourties []string
}

func main() {
	neil := person{
		first:      "Neil",
		last:       "Liddicott",
		favourties: []string{"chocolate", "vanilla"},
	}

	helen := person{
		first:      "Helen",
		last:       "Liddicott",
		favourties: []string{"chocolate", "just chocolate!"},
	}

	fmt.Println(neil)
	fmt.Print(helen)

	fmt.Println("")
	fmt.Println(neil.first)
	fmt.Println(neil.last)

	for i, v := range neil.favourties {
		fmt.Printf("Ice Cream [%v] [%v]\n", i, v)
	}
}
