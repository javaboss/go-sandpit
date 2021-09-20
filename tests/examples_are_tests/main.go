package main

// Testable Examples in Go...
// https://go.dev/blog/examples

import (
	"fmt"
	"go-sandpit/tests/doc_tests/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
