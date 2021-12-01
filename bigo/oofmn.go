package main

import "fmt"

func main() {
	mcrossn(5, 10)
}

// o(m * n)
func mcrossn(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("m : %d, n : %d \n", i, j)
		}
	}
}
