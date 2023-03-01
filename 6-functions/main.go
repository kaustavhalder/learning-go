package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the main function")

	// sa := secretAgent{
	// 	person: person{
	// 		"james",
	// 		"bond",
	// 	},
	// 	ltk: true,
	// }
	// fmt.Println(sa)

	// Function Expression
	f := func() {
		fmt.Println("this is an function expression")
	}
	f()
	foo()
	bar()

}

// Returning a string
func foo() string {
	s := "hello world"
	return s
}

// Returning a function
func bar() func() int {
	return func() int {
		return 451
	}
}