package main

import "fmt"

// Functions in GO are always pass by value and not reference
// GO is typically not oop oriented, but there are oop features

// func foo() {
// 	fmt.Println("this is called from foo")
// }

// // Paramters defined below
// func bar(s string) {
// 	fmt.Println("hello", s)
// }

// // Retrun types are defined below
// func woo(s string) string {
// 	return "Wooo"
// }

// // Retrun types are defined below
// // Different return type , match the funtion signature
// func test(fn string, ln string) (string, bool) {
// 	fmt.Println("this is from te test fucntion")
// 	return fn, true
// }

// Variadic Parameters
// Number of arguments can be zero or more
// Should always be the final parameter of the function signature
// func sum(x ...int) int {
// 	fmt.Println(x)
// 	// Transforms all the things to a slice
// 	sum := 0
// 	for i, v := range x {
// 		fmt.Println("i: ", i, "v: ", v)
// 		sum += v
// 	}
// 	return sum
// }

// Defered Functions
// Defer and execution of a function
// func foo() {
// 	fmt.Println("foo")
// }
// func bar() {
// 	fmt.Println("bar")
// }

// Methods

// Create a new struct

// type person struct {
// 	first string
// 	last  string
// }

// type secretAgent struct {
// 	person
// 	ltk bool
// }

// Below snippet is an example of a receiver attached to a function
// when a receiver is defined, it is attached to the afore mentioned struct
// and any new created struct has access to this function through the dot notation
// func (s secretAgent) speak() {
// 	fmt.Println("Hello my name is ", s.first, " ", s.last)
// }

func main() {
	fmt.Println("this is for functions")
	// foo()
	// bar("Kevin")
	// s := woo("ALALAL")
	// fmt.Println(s)
	// // Use short declaration operator within a fucntion body
	// x, y := test("Vlad", "Rodmanavich")
	// fmt.Println(x)
	// fmt.Println(y)
	// total := sum(1, 2, 3, 4, 5, 4, 4, 4, 45)
	// fmt.Println("Total is ", total)

	// // Unfurling a slice
	// xi := []int{1, 2, 4, 5, 6, 4, 55}
	// // Calling and unfurl
	// total := sum(xi...)
	// fmt.Println("Total is ", total)

	// Defered
	// defer foo()
	// bar()

	//Attaching methods to a struct
	// sa := secretAgent{
	// 	person: person{
	// 		"james",
	// 		"bond",
	// 	},
	// 	ltk: true,
	// }
	// fmt.Println(sa)
	// sa.speak()

}
