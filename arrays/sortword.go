package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
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

func sortUsingRegExp(str string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(str, -1)
}

func main() {
	var sortedString string
	m := make(map[int]string)
	m1 := make(map[int]string)
	Input := "is2 Thi1s T4est 3a"

	nums := sortUsingRegExp(Input) // This will work on any numberic string data
	for i, word := range strings.Fields(Input) {
		i, err := strconv.Atoi(nums[i])
		if err == nil {
			m1[i] = word
		}
	}
	sortedString = sortMap(m1)
	fmt.Println(sortedString)
	s := strings.Fields(Input)
	for _, word := range s {
		num, err := getNum(word)
		if err == nil {
			m[int(num)] = word
		}
	}
	sortedString = sortMap(m)
	fmt.Println(sortedString)

}

func sortMap(m map[int]string) string {
	var str string
	var keys []int
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
	return str
}
