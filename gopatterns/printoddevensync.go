package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			printEven(i)
			ch <- i
		}
	}()
	func() {
		for i := 0; i <= 10; i++ {
			printOdd(i)
			_, ok := <-ch
			if ok {
				fmt.Printf("")
			}

		}
	}()
	close(ch)
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
