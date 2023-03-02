package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the main function")

	// Function Expression
	f := func() {
		fmt.Println("this is an function expression")
	}
	f()
	foo()
	bar()
	b2 := bar2()
	fmt.Println(b2())
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

func bar2() func() string {
	return func() string {
		fmt.Println("Alibaba")
		return "Alibaba"
	}
}
