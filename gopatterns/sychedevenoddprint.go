package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)

	go func() {
		for i := 0; i <= 10; i++ {
			printEven(i)
		}
		ch <- true
	}()

	go func() {
		if <-ch {
			for i := 0; i <= 10; i++ {
				printOdd(i)
			}
		}
		ch <- true
	}()

	<-ch
	fmt.Printf("\nend of program")
}

func printEven(i int) {
	if i%2 == 0 {
		fmt.Printf("%d ", i)
	}
}

func printOdd(i int) {
	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}
}
