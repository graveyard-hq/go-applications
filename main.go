package main

import (
	"fmt"
)

var year int = 2022 // This is a variable.

// main function is the entry point of a Go program
func main() {
	fmt.Printf("Hello")
}

// Declaring multiple variables.

var (
	name    string = "John"   // String.
	age     int    = 30       // Integer.
	address string = "London" // String.
	works   bool   = true     // Bolean.

)

var foo string = "bar"

// You cannot redeclare a variable.
// But you can shadow a variable.

// All variables in go must be used.
// Variables are declared with the var keyword.

// Naming convention:
// 1. Lowercase
// 2. No spaces
// 3. No special characters
// 4. No numbers
// camelCase or PascalCase.
// Capitalize acronyms.
// As short as possible.

// Type conversion
// destunationType(variable)
// use strconv package for strings.

/*
 2.Primitive types
*/

// bool [ Boolean type aka "True" or "False" ]

// Numeric types [
// - Integers
// - Floating point
// - Complex numbers
// ]

// Text Types
// - String
// - Rune
// - Byte

// Example of a Boolean type.

func exampleBol() {
	var x bool = true // Try false too.
	fmt.Println(x)
}

// should print true.

// Lets try this -

// Go has zero values for all types.
// For example, the zero value for an int is 0.
// The zero value for a string is the empty string "".
// The zero value for a bool is false.

// The zero value for a struct is defined as all zeros.
// For example, the zero value for a struct {X int; Y int} is {0, 0}.

// The zero value for an array is nil.

// int chart
// int8 int16 int32 int64

func test() {
	var x int = 42
	fmt.Println(x)
}

//unit
// byte uint8
// rune uint32
// uint uint
// uintptr uintptr
// float32 float64

// Operators in Go
// Arithmetic operators
// - Addition
// - Subtraction
// - Multiplication
// - Division
// - Modulus
// - Exponent

// What is this =:?
// Short declaration operator.
// Learn more here https://go.dev/tour/basics/10

func calculate() {
	a := 10
	b := 5

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(a ^ b)

}

// Bit operators
// &
// |
// ^
// <<
// >>

// Bit Shift Operators
// <<
// >>
// &^

// Logical operators
// &&
// ||
// !

// If Else
