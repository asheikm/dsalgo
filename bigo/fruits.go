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

type Fillerer interface {
	filler()
}

/* O(n) --> Linear operation
   staight forward or brute force search to find something
*/

func main() {
	// Tell the system to use cpus available, let us see if it means ot anything
	// with go routines with channels and sync package for communication
	runtime.GOMAXPROCS(runtime.NumCPU())
	strslice := make([]string, 1, 1)

	// no go routines
	if NOTHREAD {
		runWithNoGo(strslice)
	}

	// check by enabling wait group
	if WGCODE {
		runWithSyncGo(strslice)
	}

	// check by using channel communication
	if CHANCODE {
		runWithChanGo(strslice)
	}

	inchan := make(chan []string)
	var wg sync.WaitGroup
	wg.Add(2)
	wgFillerReceiver := wgFiller{10000, strslice, "banana", &wg}
	chFillerReceiver := chanFiller{10000, strslice, "banana", inchan}
	FillerReceiver := Filler{10000, strslice, "banana"}

	f := []Fillerer{&FillerReceiver, &wgFillerReceiver, &chFillerReceiver}
	for _, v := range f {
		go v.filler()
	}
	go func() {
		defer wgFillerReceiver.wg.Done()
		for _, v := range wgFillerReceiver.strslice {
			if strings.Compare("banana", v) == 0 {
				// fmt.Printf("Banana found \n")
			}
		}
	}()

	wgFillerReceiver.wg.Wait()
	fmt.Printf("len of strslice - fillerer: %d\n", len(wgFillerReceiver.strslice))

	inslice := <-chFillerReceiver.chanin
	for _, v := range inslice {
		if strings.Compare("banana", v) == 0 {
			// fmt.Printf("Banana found \n")
		}
	}
	fmt.Printf("len of strslice - fillerer: %d\n", len(wgFillerReceiver.strslice))

	for _, v := range FillerReceiver.strslice {
		if strings.Compare("banana", v) == 0 {
			// fmt.Printf("Banana found \n")
		}
	}
	fmt.Printf("len of strslice - fillerer: %d\n", len(wgFillerReceiver.strslice))
}

// filler for wg
func (wgfill *wgFiller) filler() {
	defer wgfill.wg.Done()
	for i := 0; i < wgfill.n; i++ {
		wgfill.strslice = append(wgfill.strslice, wgfill.fillerStr)
	}

}

// filler for chan
func (chfill *chanFiller) filler() {
	for i := 0; i < chfill.n; i++ {
		chfill.strslice = append(chfill.strslice, chfill.fillerStr)
	}
	chfill.chanin <- chfill.strslice
}

// filler for no goroutine
func (fill *Filler) filler() {
	for i := 0; i < fill.n; i++ {
		fill.strslice = append(fill.strslice, fill.fillerStr)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s to complete\n", name, elapsed)
}

func runWithSyncGo(strslice []string) {
	defer timeTrack(time.Now(), "syncode")
	var wg sync.WaitGroup
	wgFillerReceiver := wgFiller{10000, strslice, "banana", &wg}
	wgFillerReceiver.wg.Add(2)
	go wgFillerReceiver.filler()
	go func() {
		defer wgFillerReceiver.wg.Done()
		for _, v := range wgFillerReceiver.strslice {
			if strings.Compare("banana", v) == 0 {
				// fmt.Printf("Banana found \n")
			}
		}
	}()
	wgFillerReceiver.wg.Wait()
	fmt.Printf("len of strslice: %d\n", len(wgFillerReceiver.strslice))
}

func runWithChanGo(strslice []string) {
	defer timeTrack(time.Now(), "chancode")
	inchan := make(chan []string)
	chFillerReceiver := chanFiller{10000, strslice, "banana", inchan}
	go func() {
		chFillerReceiver.filler()
	}()

	inslice := <-chFillerReceiver.chanin
	for _, v := range inslice {
		if strings.Compare("banana", v) == 0 {
			// fmt.Printf("Banana found \n")
		}
	}
	fmt.Printf("len of strslice: %d\n", len(inslice))
}

func runWithNoGo(strslice []string) {
	defer timeTrack(time.Now(), "nogoroutine")

	FillerReceiver := Filler{10000, strslice, "banana"}
	FillerReceiver.filler()

	for _, v := range FillerReceiver.strslice {
		if strings.Compare("banana", v) == 0 {
			// fmt.Printf("Banana found \n")
		}
	}
	fmt.Printf("len of strslice: %d\n", len(FillerReceiver.strslice))
}
