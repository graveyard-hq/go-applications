package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var choice int
	fmt.Println("Choose one:")
	fmt.Println("1. Rock")
	fmt.Println("2. Paper")
	fmt.Println("3. Scissors")
	fmt.Scanf("%d", &choice)

	var computerChoice int
	computerChoice = rand.Intn(3) + 1

	fmt.Println("You chose", choice)
	fmt.Println("The computer chose", computerChoice)

	var winner int
	if choice == computerChoice {
		winner = 0
	} else if choice == 1 && computerChoice == 2 {
		winner = 2
	} else if choice == 1 && computerChoice == 3 {
		winner = 1
	} else if choice == 2 && computerChoice == 1 {
		winner = 1
	} else if choice == 2 && computerChoice == 3 {
		winner = 2
	} else if choice == 3 && computerChoice == 1 {
		winner = 2
	} else if choice == 3 && computerChoice == 2 {
		winner = 1
	}

	if winner == 0 {
		fmt.Println("It's a tie!")
	} else if winner == 1 {
		fmt.Println("You won!")
	} else {
		fmt.Println("The computer wons!")
	}
}
