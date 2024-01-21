package main

import "fmt"

func countOccurrences(str, substr string) int {
	count := 0
	substrLen := len(substr)
	for i := 0; i <= len(str)-substrLen; i++ {
		if str[i:i+substrLen] == substr {
			count++
		}
	}
	return count
}

func main() {
	str := "hellohellohello"
	substr := "hello"
  
	count := countOccurrences(str, substr)

	fmt.Printf("Number of occurrences of '%s' in '%s': %d\n", substr, str, count)
}
