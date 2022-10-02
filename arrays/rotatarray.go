package main

import (
	"fmt"
)

func main() {
	inputArray := [5]int{1, 2, 3, 4, 5}
	k := 2
	// sol1
	out := rotateArray(inputArray, k)
	fmt.Println(out)
	// sol2
	rotate(inputArray, k)
	fmt.Println(inputArray)
}

// This one bit improved
func rotate(nums []int, k int)  {
    if len(nums) == 1 {
        return
    }
    if k > len(nums) {
        k = k % len(nums)
    }
    lastHalf := nums[len(nums)-k:]
    firstHalf := nums[:len(nums)-k]
    combined := append(lastHalf, firstHalf...)
    for i := range nums {
        nums[i] = combined[i]
    }
}

// This solution can be improved
func rotateArray(inputArray [5]int, pos int) []int {
	var tempArray [5]int
	var tempPos = pos
	tempArray = inputArray
	var outputArray []int
	if pos == 0 {
		for i := 0; i <= len(inputArray)-1; i++ {
			outputArray = append(outputArray, tempArray[i])
		}
		return outputArray
	}
	for i, j := 0, len(inputArray)-1; pos != 0; i++ {
		inputArray[i], inputArray[j] = inputArray[j], inputArray[i]
		j--
		pos--
	}
	tempArrayLen := len(tempArray)
	adjustPos := tempArrayLen - tempPos - 1
	outputArray = append(inputArray[0:tempPos])
	for i := 0; i <= adjustPos; i++ {
		outputArray = append(outputArray, tempArray[i])
	}
	return outputArray
}
