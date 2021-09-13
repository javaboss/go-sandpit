package main

// Instead of using the blank identifier, make sure the code is checking and handling the error.

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1) // could just sue bs, _ here and throw the error away - instead we are checking it and logging

	if err != nil { // check and log
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
