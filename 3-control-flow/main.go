package main

import "fmt"

func main() {
	fmt.Println("this is a test")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("EVEN:", i)
		} else {
			fmt.Println("ODD:", i)
		}
	}
}
