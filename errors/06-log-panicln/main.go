package main

import (
	"log"
	"os"
)

func main() {

	defer ohdear()

	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		log.Panicln(err)
		//		panic(err)
	}
}

func ohdear() {
	println("Deferred funcs are called!")
}

/*
Panicln is equivalent to Println() followed by a call to panic().
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
