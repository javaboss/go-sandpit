// Package acdc asks if you are ready to rock
package acdc

// Sum adds an unlimited number of values fo type int
func Sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}
