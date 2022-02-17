package main

import (
	"fmt"
)

var greet = "hi"

func main(){
	var greetMom string = "mom!"
	const ageIs int = 30
	
	fmt.Println("hi ", greetMom)
	greetMom = "mom"
	fmt.Println("say ", greetMom)

	// Formatting fmt aka print
	fmt.Printf("hi %v how are you???", greetMom)
	// use \n for new line
	fmt.Printf("hi %v how are you???\n", greetMom)

	// Data types >>>
	// var nameOfUser
	// nameOfUser = "John"
	// fmt.Printf(nameOfUser) // Won't work!

	// Basic one's are String & Integers
	// Other's are
	// Booleans, Maps, Arrays
	// Define a type
	var nameOfUser string
	var ageOfUser int

	nameOfUser = "John"
	ageOfUser = 30
	fmt.Printf("Age of %v is %v.\n", nameOfUser, ageOfUser)

	// print types
	fmt.Printf("type is %T", ageOfUser) // should be an integer (int)

}
