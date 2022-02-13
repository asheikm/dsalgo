package main

import (
	"fmt"
	"sort"
)

// 1,2,4,2,2,4,7,5,2,9,1  -- input
// 2,2,2,2,1,1,4,4,7,5,9  -- Frequency based [TBD]
// 1,1,2,2,2,2,4,4,5,7,9  -- ordered

var input [11]int = [11]int{1, 2, 4, 2, 2, 4, 7, 5, 2, 9, 1}

func main() {
	m := getDuplicatesCount(input)
	fmt.Println("map :", m)
	mo := sortMapByKey(m)
	o := orderedArray(mo)
	for i := 0; i < len(o); i++ {
		fmt.Printf("%d ", o[i])
	}
}

func getDuplicatesCount(n [11]int) map[int]int {
	nmap := make(map[int]int)
	count := 0
	for i := 0; i < len(n); i++ {
		if len(nmap) == 0 {
			nmap[n[i]] = count + 1
		} else {
			nmap[n[i]] = 1 + nmap[n[i]]
		}
	}
	fmt.Println(nmap)
	return nmap
}

func orderedArray(m map[int]int) []int {
	var orderdSliceArray []int = nil
	for k, v := range m {
		var t []int = nil
		orderdSliceArray = append(orderdSliceArray, fillArray(t, k, v)...)
	}
	return orderdSliceArray
}

func fillArray(sa []int, n int, c int) []int {
	for i := 0; i < c; i++ {
		sa = append(sa, n)
	}
	return sa
}

// sort map based on key
func sortMapByKey(m map[int]int) map[int]int {
	v := make([]int, 0, len(m))
	o := make(map[int]int)
	for value := range m {
		v = append(v, value)
	}

	sort.Ints(v)

	for _, value := range v {
		o[value] = m[value]
	}
	return o
}
