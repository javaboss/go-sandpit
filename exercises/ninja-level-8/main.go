package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	exercise1and2()
	exercise3()
}

func exercise1and2() {
	type person struct {
		First string
		Age   int
	}

	type unmarshalledPerson []struct {
		First string `json:"First"`
		Age   int    `json:"Age"`
	}

	// EXERCISE 1 - JSON Marshal
	u1 := person{
		First: "James",
		Age:   32,
	}

	u2 := person{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := person{
		First: "M",
		Age:   54,
	}

	users := []person{u1, u2, u3}

	fmt.Println(users)
	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("BS:", string(bs))
	}

	// EXERCISE 2 - JSON Unmarshal

	// this is using a generated slice of struct from https://mholt.github.io/json-to-go/
	// you could just do var newUsers []user

	var newUsers unmarshalledPerson
	err = json.Unmarshal(bs, &newUsers)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(newUsers)

		for i, user := range newUsers {
			fmt.Println("Person number #", i+1)
			fmt.Println(user)
		}
	}
}

func exercise3() {
	// EXERCISE 3 - JSON encode with a customer encoder to stdout
	type user struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Exercise 3")

	fmt.Println("users")
	fmt.Println(users)

	fmt.Println("Encoding to stdout")
	err := json.NewEncoder(os.Stdout).Encode(users) // using an encoder (os.stdout)
	if err != nil {
		fmt.Println("Error:", err)
	}
}