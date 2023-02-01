package main

import (
  "fmt"
)

func binarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (left + right) / 2
		if array[middle] == target {
			return middle
		} else if array[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	index := binarySearch(array, target)
	fmt.Println(index)
}
