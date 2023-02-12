package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	fmt.Println("this section is for structs")
	p1 := person{
		first: "james",
		last:  "bond",
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
}
