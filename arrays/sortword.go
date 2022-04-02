package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
Your task is to sort a given string. Each word in the String will contain a single number.
        This number is the position the word should have in the result.

        Note: Numbers can be from 1 to 9. So 1 will be the first word (not 0).

        If the input String is empty, return an empty String.
        The words in the input String will only contain valid consecutive numbers.
        For an input: "is2 Thi1s T4est 3a" the function should return "Thi1s is2 3a T4est"
*/

func getNum(word string) (byte, error) {
	for i := 0; i < len(word); i++ {
		if word[i] < '9' && word[i] > '0' { // should be used && instead of or , mistake 1
			return word[i], nil
		}
	}
	return 0, errors.New("not found")
}

func sortNumsFromStringUsingRegExp(str string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(str, -1)
}

func main() {
	var sortedString string
	m := make(map[int]string)
	m1 := make(map[int]string)
	Input := "is2 Thi1s T4est 3a"

	nums := sortNumsFromStringUsingRegExp(Input) // This will work on any numberic string data
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
	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		str += m[keys[i]] + " "
	}
	return str
}
