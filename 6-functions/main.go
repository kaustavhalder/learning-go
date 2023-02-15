package main

import "fmt"

// Functions in GO are always pass by value and not reference
// GO is typically not oop oriented, but there are oop features

func foo() {
	fmt.Println("this is called from foo")
}

func bar(s string) {
	fmt.Println("hello", s)
}

func woo(s string) string {
	return "Wooo"
}

func test(fn string, ln string) (string, bool) {
	fmt.Println("this is from te test fucntion")
	return fn, false
}
func main() {
	fmt.Println("this is for functions")
	foo()
	bar("Kevin")
	s := woo("ALALAL")
	fmt.Println(s)
	x, y := test("Vlad", "Rodmanavich")
	fmt.Println(x)
	fmt.Println(y)
}
