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
	return "hello world"
}

// Returning a function
// need to mention the return type of the inner function as well
func bar() func() int {
	return func() int {
		return 451
	}
}
