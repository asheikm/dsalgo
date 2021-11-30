package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const CHANCODE = true
const WGCODE = true
const NOTHREAD = true

type wgFiller struct {
	n         int
	strslice  []string
	fillerStr string
	wg        *sync.WaitGroup
}

type chanFiller struct {
	n         int
	strslice  []string
	fillerStr string
	chanin    chan []string
}

type Filler struct {
	n         int
	strslice  []string
	fillerStr string
}

/* O(n) --> Linear operation
   staight forward or brute force search to find something
*/

func main() {
	// Tell the system to use cpus available, let us see if it means ot anything
	// with go routines with synchronization using sync package
	runtime.GOMAXPROCS(runtime.NumCPU())
	strslice := make([]string, 1, 1)

	var wg sync.WaitGroup

	// check by enabling wait group
	if WGCODE {
		runWithSyncGo(&wg, strslice)
	}

	// check by using channel communication
	if CHANCODE {
		runWithChanGo(strslice)
	}

	// no go routines
	if NOTHREAD {
		runWithNoGo(strslice)
	}
}

// filler for wg
func (wgfill *wgFiller) filler() []string {
	for i := 0; i < wgfill.n; i++ {
		wgfill.strslice = append(wgfill.strslice, wgfill.fillerStr)
	}
	wgfill.wg.Done()
	return wgfill.strslice
}

// filler for chan
func (chfill *chanFiller) filler() {
	for i := 0; i < chfill.n; i++ {
		chfill.strslice = append(chfill.strslice, chfill.fillerStr)
	}
	chfill.chanin <- chfill.strslice
}

// filler for no goroutine
func (fill *Filler) filler() []string {
	for i := 0; i < fill.n; i++ {
		fill.strslice = append(fill.strslice, fill.fillerStr)
	}
	return fill.strslice
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s to complete\n", name, elapsed)
}

func runWithSyncGo(wg *sync.WaitGroup, strslice []string) {
	strArray := [5]string{"apple", "orange", "pinapple", "mango", "banana"}
	defer timeTrack(time.Now(), "syncode")
	wg.Add(3)
	go func() {
		wgFillerMethod := wgFiller{10000, strslice, "banana", wg}
		strslice = wgFillerMethod.filler()
	}()
	go func() {
		for _, v := range strArray {
			if strings.Compare("mango", v) == 0 {
				// fmt.Printf("Banana found \n")
			}
		}
		wg.Done()
	}()

	go func() {
		for _, v := range strslice {
			if strings.Compare("banana", v) == 0 {
				// fmt.Printf("Banana found \n")
			}
		}
		wg.Done()
	}()
	fmt.Printf("len of strslice: %d\n", len(strslice))
	wg.Wait()
}

func runWithChanGo(strslice []string) {
	defer timeTrack(time.Now(), "chancode")
	inchan := make(chan []string)
	go func() {
		chFillerMethod := chanFiller{10000, strslice, "banana", inchan}
		chFillerMethod.filler()
	}()

	inslice := <-inchan
	for _, v := range inslice {
		if strings.Compare("banana", v) == 0 {
			// fmt.Printf("Banana found \n")
		}
	}
	fmt.Printf("len of strslice: %d\n", len(strslice))
}

func runWithNoGo(strslice []string) {
	defer timeTrack(time.Now(), "nogoroutine")

	FillerMethod := Filler{10000, strslice, "banana"}
	slicestr := FillerMethod.filler()

	for _, v := range slicestr {
		if strings.Compare("banana", v) == 0 {
			// fmt.Printf("Banana found \n")
		}
	}
	fmt.Printf("len of strslice: %d\n", len(strslice))
}

/**************************************************************************
Sample output:
--------------
D:\DEV\ds\dsalgo\bigo>go run fruits.go
len of strslice: 10001
Took  505.5µs to complete the sync code
len of strslice: 10001
Took  882.4µs to complete the chan code
len of strslice: 10001
Took  525.4µs to complete without chan or sync code

D:\DEV\ds\dsalgo\bigo>go run fruits.go
len of strslice: 10001
Took  0s to complete the sync code
len of strslice: 10001
Took  340.9µs to complete the chan code
len of strslice: 10001
Took  480.8µs to complete without chan or sync code

D:\DEV\ds\dsalgo\bigo>go run fruits.go
len of strslice: 10001
Took  491.2µs to complete the sync code
len of strslice: 10001
Took  0s to complete the chan code
len of strslice: 10001
Took  498.6µs to complete without chan or sync code

D:\DEV\ds\dsalgo\bigo>go run fruits.go
# command-line-arguments
.\fruits.go:110:2: undefined: wg
.\fruits.go:111:20: undefined: start

D:\DEV\ds\dsalgo\bigo>go run fruits.go
len of strslice: 10001
Took  520µs to complete the chan code
len of strslice: 10001
Took  605.7µs to complete without chan or sync code
len of strslice: 10001
Took  1.6844ms to complete the sync code
*************************************************************************/
