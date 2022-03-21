package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.globalsign.com/en"
	word := "globalsign"
	count, err := getGivenWordCountFromUrl(url, word)
	if err == nil {
		fmt.Printf("Number of given word %s in given url %s is : %d\n", word, url, count)
	}
	return
}

func getGivenWordCountFromUrl(url, word string) (any, error) {
	count := 0
	b, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer b.Body.Close()

	var bd []byte

	bd, err = ioutil.ReadAll(b.Body)
	if err != nil {
		fmt.Errorf("error", err)
		return nil, err
	}
	words := strings.Fields(string(bd))
	for _, cword := range words {
		if strings.Contains(cword, word) {
			count++
		}
	}
	return count, nil
}
