package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 0}
	c := getSum(a) // o(n)
	d := getSum(b) // o(m) or o(n)
	// o(n) + o(n) ==> o(2n) ==> o(n)
}

// o(n)
func getSum(m [5]int) {
	var n int
	for i, v := range m {
		n += v
	}
	return n
}
