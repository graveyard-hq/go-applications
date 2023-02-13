package main

import (
  "fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FibonacciSearch(arr []int, x int) int {
	n := len(arr)
	var fib2 = 0
	var fib1 = 1
	var fib = fib1 + fib2
	for fib < n {
		fib2 = fib1
		fib1 = fib
		fib = fib1 + fib2
	}
	var offset = -1
	for fib > 1 {
		var i = min(offset + fib2, n - 1)
		if arr[i] < x {
			fib = fib1
			fib1 = fib2
			fib2 = fib - fib1
			offset = i
		} else if arr[i] > x {
			fib = fib2
			fib1 = fib1 - fib2
			fib2 = fib - fib1
		} else {
			return i
		}
	}
	if fib1 == 1 && arr[offset+1] == x {
		return offset + 1
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := FibonacciSearch(arr, 6)
	fmt.Println("Index of 6:", result)
}
