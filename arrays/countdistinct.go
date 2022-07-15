/*
Given an array of size n and an integer k, return the count of distinct numbers in all windows of size k.
Input: arr[] = {1, 2, 1, 3, 4, 2, 3};
       k = 4
Output: 3 4 4 3

Explanation:
First window is {1, 2, 1, 3}, count of distinct numbers is 3
Second window is {2, 1, 3, 4} count of distinct numbers is 4
Third window is {1, 3, 4, 2} count of distinct numbers is 4
Fourth window is {3, 4, 2, 3} count of distinct numbers is 3

Input: arr[] = {1, 2, 4, 4};
       k = 2
Output: 2 2 1
*/
package main

import "fmt"

func main() {
	intputArray := []int{1, 2, 4, 4, 4}
	windowSize := 2
	outputArray := make([]int, 0, 1)
	// window size check
	arrayLen := len(intputArray)
	if arrayLen <= 1 {
		fmt.Println("Invalid array size")
		return
	}
	if windowSize == 0 || windowSize > arrayLen {
		fmt.Println("Invalid window size")
		return
	}
	for i := 0; i < arrayLen; i++ {
		curSlicerange := windowSize + i
		if curSlicerange <= arrayLen {
			newWindowedArray := intputArray[i:curSlicerange]
			rc := 0
			for i := 0; i < windowSize; i++ {
				for j := i + 1; j < (windowSize - i); j++ {
					if newWindowedArray[i] != newWindowedArray[j] {
						rc = rc + 1
					}
				}
			}
			if windowSize == 2 {
				rc = rc + 1
			}
			outputArray = append(outputArray, rc)
		}
	}
	fmt.Println(outputArray[:])
}
