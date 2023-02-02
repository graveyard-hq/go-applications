package main

import (
  "fmt"
)

func linearSearch(arr []int, key int) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}
	key := 110
	result := linearSearch(arr, key)
	if result == -1 {
		fmt.Println("Element is not present in array")
	} else {
		fmt.Printf("Element is present at index %d\n", result)
	}
}
