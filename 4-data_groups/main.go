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
	// for i, v := range x {
	// 	fmt.Println("Index", i, v)
	// }
	// fmt.Println(x[0:1])
	// fmt.Println(x[1:])
	// for i := 0; i < len(x); i++ {
	// 	fmt.Println("slice from for loop")
	// 	fmt.Println(x[i])
	// }
	x = append(x, 44)
	fmt.Println("this is the slice after 44", x)
	y := []int{98, 87, 65, 53}
	x = append(x, y...) // this is sort of an unpacking operation being done on the slice and or any datatypes
	fmt.Println("this is the slice after appending another slice", x)
	// Delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println("slice after the delete operation", x)
	// Make - using this to crate a slice make([]int, 10, 1000)

	// Maps
	m := map[string]int{
		"James": 32,
		"Novak": 27,
		"Roger": 33,
	}
	n := map[int]string{
		1: "Vlad",
		2: "Nikolai",
		3: "Aleksi",
	}
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Roger"]) // Zero value is returned
	v, ok := m["Roger"]
	fmt.Println(v)
	fmt.Println(ok)
	// comma, ok syntax 
	if v, ok := m["Roger"]; ok {
		fmt.Print("PRint ", v)
	}
	m["Andrei"] = 44
	fmt.Println(m)
	n[4] = "Whiskey"
	// fmt.Println(n)

}
