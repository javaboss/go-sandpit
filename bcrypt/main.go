package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password1234"

	fmt.Println("PRE-ENCRYPT: ", s)

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("There was an error", err)
	} else {
		fmt.Println("POST-ENCRYPT: ", bs)
	}
}
