package main

import "fmt"

type person struct {
	first      string
	last       string
	favourties []string
}

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

func main() {
	// EXERCISE 1
	fmt.Println("EXERCISE 1")

	neil := person{
		first:      "Neil",
		last:       "Liddicott",
		favourties: []string{"chocolate", "vanilla"},
	}

	helen := person{
		first:      "Helen",
		last:       "Keys",
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

	// EXERCISE 2
	fmt.Println("EXERCISE 2")
	m := map[string]person{
		neil.last:  neil,
		helen.last: helen,
	}

	fmt.Println(m)

	// range over map
	for k, v := range m {
		//		fmt.Printf("KEY [%v]\tVALUE[%v]\n", k, v)
		fmt.Println("KEY:", k)
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.favourties {
			fmt.Println(i, val)
		}
	}

	// EXERCISE 3
	fmt.Println("EXERCISE 3")
	neilsTruck := truck{
		vehicle: vehicle{
			doors:  4,
			colour: "black",
		},
		fourWheel: true,
	}

	fmt.Println(neilsTruck)
}
