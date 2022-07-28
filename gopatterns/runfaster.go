/*
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		printData(&i)
	}
}

func printData(i *int) {
	fmt.Println(*i)
	time.Sleep(1 * time.Second)
}

*/

// Fix the above code to run faster without removing time.Sleep should also print 1 to 10

/* package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printData(i, &wg)
	}
	wg.Wait()
}

func printData(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	time.Sleep(1 * time.Second)
	wg.Done()
}
*/

// Above program runs 1 -10 not in order , fix that code to run them 1..2..3.. 10
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
		wg.Add(1)
		go printData(&wg, ch)
	}
	wg.Wait()

}

func printData(wg *sync.WaitGroup, ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println(<-ch)
	wg.Done()
}
