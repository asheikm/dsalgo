package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is test"
	strlen := len(s)
	var temp []byte
	var sr string
	j := 0
	// reverse words in given string
	for i := 0; i < strlen; i++ {
		if strings.Compare(string(s[i]), " ") != 0 {
			temp = append(temp, byte(s[i]))
			j++
		} else {
			sr += reverse(string(temp)) + " "
			temp = []byte{}
		}
	}
	sr += reverse(string(temp)) + " "
	fmt.Println(sr)
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = s[j], s[i]
	}
	return string(r)
}
