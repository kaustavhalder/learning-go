package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("this is the exercise ")
	// 1
	// x := 42
	// y := "James Bond"
	// z := true
	// fmt.Println(x, y, z)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// 2 zero values
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	fmt.Println(s)

}
