package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

type unmarshalledUser []struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	// EXERCISE 1 - JSON Marshal
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

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

	var newUsers unmarshalledUser
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
