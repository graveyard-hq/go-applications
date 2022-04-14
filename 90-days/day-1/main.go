package main

import "fmt"

type Person struct {
	name    string
	age     int
	country string
}

func newUser(name string, age int, country string) *Person {
	newPer := Person{name: name}
	newPer = Person{age: age}
	newPer = Person{country: country}
	return &newPer
}

func main() {
	fmt.Println(Person{
		name:    "Aelpxy",
		age:     15,
		country: "null",
	})
}
