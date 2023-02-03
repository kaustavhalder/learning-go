package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("EVEN:", i)
	// 	} else {
	// 		fmt.Println("ODD:", i)
	// 	}
	// }
	switch {
	case false:
		fmt.Println("this is false")
	case true:
		fmt.Println("this is true")
	default:
		fmt.Println("this is default value")
	}
}
