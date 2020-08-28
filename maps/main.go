package main

import "fmt"

func main() {
	fmt.Println("Hello")

	m := map[string]int{
		"james":      36,
		"moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["james"])
	fmt.Println(m["moneypenny"])

	fmt.Println(m["henry"])

	// checking if key / value exists
	v, ok := m["henry"]
	fmt.Println(v, ok)

	if v, ok := m["henry"]; !ok {
		fmt.Println("Ouch!")
	} else {
		fmt.Println(v)
	}

	// adding a key/value
	m["Neil"] = 46

	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete key/value from map
	fmt.Println("*** DELETE ***")
	delete(m, "Neil")

	for k, v := range m {
		fmt.Println(k, v)
	}

}
