package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func getNum(word string) (byte, error) {
	for i := 0; i < len(word); i++ {
		if word[i] < '9' && word[i] > '0' { // should be used && instead of or , mistake 1
			return word[i], nil
		}
	}
	return 0, errors.New("not found")
}

func main() {
	m := make(map[byte]string)
	Input := "is2 Thi1s T4est 3a"
	s := strings.Fields(Input)
	for _, word := range s {
		num, err := getNum(word)
		if err == nil {
			m[num] = word
		}
	}
	var str string
	var keys []byte
	// Maps are not by default sorted so, i wanted to sort the keys from the map
	// first then use the keys to access the values to append and return
	for k := range m {
		keys = append(keys, k)
	}
	// I thought of default sort of Bytes available, but somehow since it is a slice we have use callback for comparision
	// no straight forward method availbale to sort Byte array
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for i := 0; i < len(keys); i++ {
		str += m[keys[i]] + " "
	}
	fmt.Println(str)
}
