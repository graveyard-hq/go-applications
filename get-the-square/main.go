package main

import "fmt"

func main() {
	var userInput int
	fmt.Println("Get the square!")
	fmt.Println("Enter the number")
	fmt.Scan(&userInput)
	var doJon = multiply(userInput, userInput)
	fmt.Printf("Square of %v is %v \n", userInput, doJon)
}

func multiply(first int, second int) int {
	return first * second
}
