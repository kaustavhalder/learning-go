package main

import "fmt"

func main() {
	fmt.Println("this is for control flow")
	for i := 0; i < 10; i++ {
		fmt.Println("Value of i is ", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j is ", j)
		}
	}
}
