package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	// Tell the system to use cpus available, let us see if it means ot anything
	// without using goroutines
	runtime.GOMAXPROCS(runtime.NumCPU())

	strslice := make([]string, 1, 1)
	start := time.Now()
	strArray := [5]string{"apple", "orange", "pinapple", "mango", "banana"}

	strslice = filler(10000, strslice, "banana")

	for _, v := range strArray {
		if strings.Compare("mango", v) == 0 {
			fmt.Printf("Banana found \n")
		}
	}

	for _, v := range strslice {
		if strings.Compare("banana", v) == 0 {
			fmt.Printf("Banana found \n")
		}
	}
	end := time.Since(start)
	fmt.Printf("len of strslice: %d\n", len(strslice))
	fmt.Println("Took ", end, "to complete")
}

// filler
func filler(n int, strslice []string, fillerstr string) []string {
	for i := 0; i < n; i++ {
		strslice = append(strslice, fillerstr)
	}
	return strslice
}
