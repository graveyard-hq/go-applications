package main

import (
	"fmt"
	"sync"
)

func generateNumbers(numbers chan<- int, done chan<- bool) {
	defer close(numbers)
	for i := 0; i < 1000; i++ {
		numbers <- i
	}
	done <- true
}

func squareNumbers(numbers <-chan int, squares chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range numbers {
		squares <- i * i
	}
}

func printSquares(squares <-chan int, done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case square, ok := <-squares:
			if ok {
				fmt.Println(square)
			} else {
				squares = nil
			}
		case <-done:
			if squares == nil {
				return
			}
		}
	}
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go printSquares(squares, done, &wg)

	wg.Add(1)
	go squareNumbers(numbers, squares, &wg)

	go generateNumbers(numbers, done)

	wg.Wait()
}
