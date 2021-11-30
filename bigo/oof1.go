package main

import (
	"fmt"
)

/* whatever the length of the array may be we would print the fist two elements
   we perform only one step to get the data and print, the notation is constant aloways
*/

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 9, 9, 10}
	printFistTwoElements(arr)

}

func printFistTwoElements(arr [10]int) {
	fmt.Println(arr[0])
	fmt.Println(arr[1])
}
