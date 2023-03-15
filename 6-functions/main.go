// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("this is the main function")

// 	// Function Expression
// 	f := func() {
// 		fmt.Println("this is an function expression")
// 	}
// 	f()
// 	foo()
// 	bar()
// 	b2 := bar2()
// 	fmt.Println(b2())
// 	ii := []int{2, 5, 5, 5}
// 	su := sum(ii...)
// 	fmt.Println(su)

// }

// // Returning a string
// func foo() string {
// 	return "hello world"
// }

// // Returning a function
// // need to mention the return type of the inner function as well
// func bar() func() int {
// 	return func() int {
// 		return 451
// 	}
// }

// func bar2() func() string {
// 	return func() string {
// 		fmt.Println("Alibaba")
// 		return "Alibaba"
// 	}
// }

// // Callback
// // Passing a function as a parameter eg: JavaScript

// func sum(x ...int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }

package main

import "fmt"

// Single return type
func printHw() string {
	fmt.Println("Hello World")
	return "hello world"
}

func baz(s string) {
	fmt.Println("hello", s)
}

// Returning multiple return types
func FullName(fn string, ln string) (string, bool) {
	fullname := fn + " " + ln
	return fullname, true
}

// Vairadic params
// ... means number of args is not given in function definition
// turns them into a slice

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		fmt.Println("itmem", i, "value ", v)
		sum += v
	}
	fmt.Println("Sum is ", sum)

}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	printHw()
	baz("Vlad")
	name, status := FullName("Ian", "Flemming")
	fmt.Println(name)
	fmt.Println(status)
	foo(2, 3, 4)
}
