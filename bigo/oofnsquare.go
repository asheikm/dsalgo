package main

import "fmt"

/*
Thank you wiki!
Multiplying two n-digit numbers by a simple algorithm; simple sorting algorithms,
such as bubble sort, selection sort, and insertion sort; (worst case) bound on
some usually faster sorting algorithms such as quicksort, Shell sort, and tree sort
*/
func main() {
	arr := []int{4, 27, 6, 18, 9, 42, 21, 90, 1}
	arrB := bubbleSort(arr)
	arrS := selectionSort(arr)
	arrI := insertionSort(arr)
}

// o(n^2) bubble sort
func bubbleSort(arr []int) []int {
	n = len(arr)
	for i := 0; i < n-1; i++ {
		// Last i elements are already in place
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1] // swap
			}
		}
	}
	return arr
}

// o(n^2) Selection sort
func selectionSort(arr []int) []int {
	// One by one move boundary of unsorted subarray
	for i := 0; i < n-1; i++ {
		// Find the minimum element in unsorted array
		min_idx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		// Swap the found minimum element with the first element
		arr[min_idx], arr[i] = arr[i], arr[min_idx]
	}
	return arr
}

// o(n^2) Insertion sort
func insertionSort(int []arr) []arr {
	n := len(arr)
	j := 0
	for i := 1; i < n; i++ {
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}
