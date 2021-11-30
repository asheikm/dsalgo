package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	strslice := make([]string, 1, 1)
	start := time.Now()
	strArray := [5]string{"apple", "orange", "pinapple", "mango", "banana"}

	filler(10000000, strslice, "banana")

	for _, v := range strArray {
		if strings.Compare("mango", v) == 0 {
			fmt.Printf("Banana found \n")
		}
	}
	end := time.Since(start)
	fmt.Println("Took ", end, "to complete")
}

// filler
func filler(n int, strslice []string, fillerstr string) {
	for i := 0; i < n; i++ {
		strslice = append(strslice, fillerstr)
	}
}
