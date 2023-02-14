package main

import "fmt"

type person struct {
	first string
	last  string
}

// Embedded types
// Similar to other programming paradigms in

type secretAgent struct {
	person
	ltk bool
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

	sAgent := secretAgent{
		person: person{
			first: "Vlad",
			last:  "Funky",
		},
		ltk: true,
	}
	fmt.Println(sAgent)

}
