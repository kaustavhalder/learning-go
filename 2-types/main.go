package main

import "fmt"

// Zero value for boolean is false

var x bool

func main() {
	fmt.Println("this package is for types")
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 42
	b := 33

	fmt.Println(a == b)
}
