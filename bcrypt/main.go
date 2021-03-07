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
	}

	fmt.Println("COMPARE ENCRYPTED PASSWORD WITH CANDIDATE PASSWORD")

	newPassword := "password1234"
	err = bcrypt.CompareHashAndPassword(bs, []byte(newPassword))

	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("PASSWORDS MATCH")
}
