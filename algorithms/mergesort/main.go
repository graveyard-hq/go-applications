package main

import (
  "fmt"
)

func mergeSort(array []int) []int {
    if len(array) <= 1 {
        return array
    }

    middle := len(array) / 2
    left := mergeSort(array[:middle])
    right := mergeSort(array[middle:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}

func main() {
    array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    sortedArray := mergeSort(array)
    fmt.Println(sortedArray)
}
