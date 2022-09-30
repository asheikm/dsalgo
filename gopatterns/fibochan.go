package main

import (
	"fmt"
)

// 0 ,1 ,1, 2, 3,5, 8, 13, 21
func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case <-quit:
			return
		case ch <- x:
			z := x + y
			x = y
			y = z
		}
	}

}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch) // read from channel ch
		}
		quit <- false
	}(n)
	fibonacci(ch, quit)
}
