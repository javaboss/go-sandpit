package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	var newPeople []person
	err = json.Unmarshal(bs, &newPeople)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newPeople)

		for i, v := range newPeople {
			fmt.Println("Person Numnber", i)
			fmt.Println(v.First, v.Last, v.Age)
		}
	}
}
