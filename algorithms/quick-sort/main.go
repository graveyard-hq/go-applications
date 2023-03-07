package main

import (
	"fmt"
)

// sorts an integer slice using the quicksort algorithm.
func Quicksort(arr []int) ([]int, error) {
	if len(arr) < 2 {
		return arr, nil
	}
	pivot := arr[0]
	var less []int
	var greater []int
	for _, value := range arr[1:] {
		if value <= pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}
	sortedLess, err := Quicksort(less)
	if err != nil {
		return nil, err
	}
	sortedGreater, err := Quicksort(greater)
	if err != nil {
		return nil, err
	}
	return append(append(sortedLess, pivot), sortedGreater...), nil
}

func main() {
	arr := []int{8, 3, 2, 5, 1, 4, 10, 7, 9, 6}
	sortedArr, err := Quicksort(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(sortedArr)
}
