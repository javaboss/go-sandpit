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
	vehicle   // no need to qualify i.e. typeOfVehicle vehicle - you can just put the struct on it's own
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
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
	helensTruck := truck{
		vehicle: vehicle{
			doors:  5,
			colour: "Carbon Black",
		},
		fourWheel: true,
	}

	neilsSedan := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "Silver",
		},
		luxury: true,
	}

	fmt.Println(helensTruck)
	fmt.Println(neilsSedan)

	fmt.Println(helensTruck.colour)
	fmt.Println(neilsSedan.luxury)

	// Exercise 4
	fmt.Println("EXERCISE 4")
	// Create an anonymous struct - include things like maps + slices inside

	as := struct {
		firstname string
		lastname  string
		friends   map[string]int
		drinks    []string
	}{
		firstname: "neil",
		lastname:  "liddicott",
		friends: map[string]int{
			"Helen": 123,
			"Jack":  456,
			"Anna":  789,
		},
		drinks: []string{
			"Stella",
			"Water",
			"Milk",
		},
	}

	fmt.Println(as)
}
