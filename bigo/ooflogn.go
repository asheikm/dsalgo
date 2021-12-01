package main

import "fmt"

func main() {
	fmt.printf("Code to demonstrate o(logn)")
	count := skiponelevel(10)
}

// o(logn)
func skiponelevel(n) {
	count := 0
	for i := 1; i < n; i = i * 2 {
		count++
	}
	return count
}
