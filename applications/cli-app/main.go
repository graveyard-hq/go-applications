package main

import "fmt"

var (
  points int = 0
)

func main() {
  askQuestion("how are you? (true/false) \n")
  checkLogic(true)

  returnPoint()
}

func askQuestion(ques string) {
  fmt.Println(ques)
}

func checkLogic(whatIs bool) {
  var storeAnswer bool;

  fmt.Scan(&storeAnswer)

  switch storeAnswer {
  case true:
    addPoints(1)
    fmt.Println("Great to hear! You just won 1 point!")
  case false:
    removePoints(1)
    fmt.Println("Sadly, you've lost 1 point.")
  default:
    removePoints(5)
    fmt.Println("Something didn't end well.")
  }
}

func returnPoint() int {
  return points
}

func addPoints(score int) int {
  points = points + 1
  return points
}

func removePoints(score int) int {
  points = points - 1
  return points
}
