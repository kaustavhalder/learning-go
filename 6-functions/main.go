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
func foo(x ...int) {
	fmt.Println(x)
	// Transforms all the things to a slice
}

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
	foo(1,2,3,4,5,4,4,4,45)
}
