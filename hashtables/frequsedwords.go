package main

/* Given the attached text file as an argument, your program will read the file, and output the 20 most frequently
   used words in the file in order, along with their frequency.
   The output should be the same to that of the following bash commands,
   where the first argument is the text file to process As mentioned,
   candidate needs to read the file and output the result of 20 most frequently used words in order and their frequency.
*/

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Filename and max freq hardcoded
	mostRepeatedTwenty := 20
	filename := "mobydick.txt"
	m := make(map[string]int)
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read file")
		return
	}
	// After reading the file as bytes convert that to string and split using space.
	// returns a slice of words each as string and insert into map where repeated frequenty will be increamented
	// on each repeation
	words := strings.Split(string(b), " ")
	for _, word := range words {
		if m[word] == 0 || len(m) == 0 {
			m[word] = 1
		} else if m[word] >= 1 {
			m[word] = m[word] + 1
		}
	}
	// Sort the map by putting the word and frequency into a struct which will be easier to sort and get required data
	sorted := sortByValue(m)
	for i := 0; i < mostRepeatedTwenty; i++ {
		fmt.Printf("word: %s, noofoccurances: %d\n", sorted[i].word, sorted[i].num)
	}

}

type kv struct {
	word string
	num  int
}

// Ref : https://go.dev/play/p/y1_WBENH4N
// https://stackoverflow.com/questions/18695346/how-can-i-sort-a-mapstringint-by-its-values#There's%20a%20new%20sort.Slice%20function
func sortByValue(m map[string]int) []kv {
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].num > ss[j].num
	})
	return ss
}

// How to run the code ?
// go run parse.go
