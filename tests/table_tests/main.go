package main

import "fmt"

// Tests must:
// 		be in a file that ends with “_test.go”
// 		put the file in the same package as the one being tested
// 		be in a func with a signature “func TestXxx(*testing.T)”
// Run a test:
// 		go test
// Deal with test failure:
// 		use t.Error to signal failure

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 7 = ", mySum(4, 7))
	fmt.Println("5 + 9 = ", mySum(5, 9))
}

func mySum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}
