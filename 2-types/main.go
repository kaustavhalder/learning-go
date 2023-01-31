package main

import (
	"fmt"
)

// Zero value for boolean is false

var x bool

// String Type
var name string

func main() {
	fmt.Println("this package is for types")
	fmt.Println(x)
	x = true
	fmt.Println(x)
	name = "Vlad"
	fmt.Println(name)
	fmt.Println(name[0])
}
