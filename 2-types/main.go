package main

import "fmt"

// Zero value for boolean is false

var x bool

func main() {
	fmt.Println("this package is for types")
	fmt.Println(x)
	x = true
	fmt.Println(x)
}
