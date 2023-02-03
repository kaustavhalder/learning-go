package main

import (
	"fmt"
)

func main() {
	fmt.Println("this section is for arrays and slices and structs and maps and all other stuff")

	// Arrays
	// var x [5]int
	// fmt.Println(x)
	// x[0] = 9
	// fmt.Println(x)
	// fmt.Println(len(x))

	// Slice
	// x := type {values} composite literal
	x := []int{2, 4, 8, 10, 12, 22, 33}
	fmt.Println("this is the slice", x)
	for i, v := range x {
		fmt.Println("Index", i, v)
	}
	fmt.Println(x[0:1])
	fmt.Println(x[1:])
	for i := 0; i < len(x); i++ {
		fmt.Println("slice from for loop")
		fmt.Println(x[i])
	}
}
