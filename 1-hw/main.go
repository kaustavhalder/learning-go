package main

import "fmt"

// when declaring outside, need to use var, short declartion uses only within function body
// declaration syntax = var {name} then type, unlike other prgramming languages
// go can implicitly identify the type of the variable, if no type is mentioned during declaration
var z int
var name string = "VLAD"
// Golang types
// strongly/static typed, so variables with declared type, cannot be used for other types 

func foo() {
	fmt.Println("this is from the foo functino")
	fmt.Println("var name is ", name)
}

func main() {
	// declare and assign at the same time
	// short declaration can only be used inside a function
	x := 43
	fmt.Println("this is a test ")
	foo()
	fmt.Println("end of program")
	x = 50
	fmt.Println(x)
	// during assignemant you can use expressions as well
	y := 100 + 24
	fmt.Println(y)
	// standard declartion of variables
	fmt.Println(z)
}
