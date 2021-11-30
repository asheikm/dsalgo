package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	// Tell the system to use cpus available, let us see if it means ot anything
	// with go routines ans synchronization using sync package
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup

	strslice := make([]string, 1, 1)
	start := time.Now()
	strArray := [5]string{"apple", "orange", "pinapple", "mango", "banana"}

	wg.Add(3)

	strslice = filler(1000000, strslice, "banana", &wg)
	go func() {
		for _, v := range strArray {
			if strings.Compare("mango", v) == 0 {
				fmt.Printf("Banana found \n")
			}
		}
		wg.Done()
	}()

	go func() {
		for _, v := range strslice {
			if strings.Compare("banana", v) == 0 {
				fmt.Printf("Banana found \n")
			}
		}
		wg.Done()
	}()

	wg.Wait()
	end := time.Since(start)
	fmt.Printf("len of strslice: %d\n", len(strslice))
	fmt.Println("Took ", end, "to complete")
}

// filler
func filler(n int, strslice []string, fillerstr string, wg *sync.WaitGroup) []string {
	for i := 0; i < n; i++ {
		strslice = append(strslice, fillerstr)
	}
	wg.Done()
	return strslice
}
