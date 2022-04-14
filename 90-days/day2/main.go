package main

import (
  "fmt"
)

var point int = 0

func add(n int) int {
  return n + point
}

func main() {
  s := make([]int, 5)

  s[0] = add(5)
  s[1] = add(10)

  i := 1
  s[3] = i
  s[4] = add(100)
  for i < 3 {
    add(1)
    fmt.Println(s)
  }
}
