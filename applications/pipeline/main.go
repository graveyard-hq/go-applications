package main

import (
        "fmt"
        "sync"
)

func generator(numbers chan<- int) {
        for i := 0; i < 1000; i++ {
                numbers <- i
        }
        close(numbers)
}

func squarer(numbers <-chan int, squares chan<- int) {
        for i := range numbers {
                squares <- i * i
        }
        close(squares)
}

func printer(squares <-chan int, wg *sync.WaitGroup) {
        for i := range squares {
                fmt.Println(i)
        }
        wg.Done()
}

func main() {
        var wg sync.WaitGroup
        numbers := make(chan int)
        squares := make(chan int)

        wg.Add(1)
        go printer(squares, &wg)

        go squarer(numbers, squares)

        generator(numbers)

        wg.Wait()
}
