package main

import "fmt"


func remove(a int) int {
  return a * 2
}

var comp string = "a"

func main() {
  s := make([]int, 10)

  s[0] = 5
  s[1] = 1
  s[2] = 0
  s[4] = 5
  s[5] = 5
  s[6] = remove(2)
  fmt.Println(s)

  i := 1
  for i < 5 {
    fmt.Println(comp)
    break
  }
}
