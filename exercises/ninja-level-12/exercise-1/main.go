package main

import (
	"fmt"
	"go-sandpit/exercises/ninja-level-12/exercise-1/dog"
)

func main() {
	// use godoc -http :8080 to run a local godoc server and see the dog documentation

	fmt.Printf("10 humnan years is %v dog years", dog.Years(10))
}
