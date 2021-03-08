package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	exercise1and2()
	exercise3()
	exercise4()
	exercise5()
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

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func exercise3() {
	// EXERCISE 3 - JSON encode with a customer encoder to stdout
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

func exercise4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLast []user

func (bl byLast) Len() int           { return len(bl) }
func (bl byLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl byLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func (u user) String() string {
	return fmt.Sprintf("%s %s:%d", u.First, u.Last, u.Age)
}

func exercise5() {

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

	fmt.Println("EXERCISE 5")

	fmt.Println("---------- USERS ----------")
	printUsers(users)

	fmt.Println("---------- SORT BY AGE ----------")
	sort.Sort(byAge(users))
	printUsers(users)

	fmt.Println(" SORT BY LAST NAME ----------")
	sort.Sort(byLast(users))
	printUsers(users)
}

func printUsers(users []user) {
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}
