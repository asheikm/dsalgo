/* package main

import (
	"fmt"
)

func main() {
	m := map[int]int{}
	ch := make(chan int)
	go readLoop(m, ch)
	writeLoop(m, ch)
}

func writeLoop(m map[int]int, ch chan int) {
	for {
		for i := 0; i < 100; i++ {
			m[i] = i
			ch <- i
		}
	}
}

func readLoop(m map[int]int, ch chan int) {
	for {
		for k, v := range m {
			fmt.Println(k, "-", v)
			<-ch
		}
	}
}
*/
package main

import (
	"fmt"
)

func main() {
	m := map[int]int{}
	boolch := make(chan bool)
	go readLoop(m, boolch)
	writeLoop(m, boolch)
}

func writeLoop(m map[int]int, ch chan bool) {
	for {
		for i := 0; i < 100; i++ {
			m[i] = i
		}
		ch <- true
	}
}

func readLoop(m map[int]int, ch chan bool) {
	for {
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		<-ch
	}
}
