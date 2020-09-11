package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := append(x[:5])
	b := append(x[5:])
	c := append(x[2:7])
	d := append(x[1:6])

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println(x)

	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// 42, 43, 44, 48, 49, 50, 51

	zz := append(z[:3], z[6:]...)
	fmt.Println(zz)

	states := make([]string, 3, 3)
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	states = []string{"New York", "Florida", "Ohio"}
	fmt.Println(states)

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	fmt.Println("SLICE OF SLICE")
	sliceofaslice := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooooo, James"}}

	for i, v := range sliceofaslice {
		fmt.Println(i, v)
	}

	// with a range it always returns index and value
	// you can discard the index using a _ like in Scala
	// for example :

	for _, v := range sliceofaslice {
		fmt.Printf("Value [%v]\n ", v)
	}

	/* MAPS */
	fmt.Println("MAPS")

	m := map[string][]string{
		"bond_james":      []string{"shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println("This is the record for: ", k)

		// note you could also use fmt.Println(i, rv)
		for i, rv := range v {
			fmt.Printf("Index [%v] Value [%v]\n", i, rv)
		}
	}

	m["neil liddicott"] = []string{"stella", "more stells", "brunettes"}

	fmt.Println("DELETE")
	delete(m, "bond_james")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
