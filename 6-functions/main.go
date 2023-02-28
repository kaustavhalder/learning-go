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

	f := func() {
		fmt.Println("this is an function expression")
	}
	f()
}
