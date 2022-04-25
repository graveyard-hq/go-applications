package main

import (
	"log"
	"fmt"
)

func main() {
	log.Println("Test CLI")
	var name string

	log.Println("What is your name?")
	fmt.Scan(&name)
	
	log.Println(name)
}
