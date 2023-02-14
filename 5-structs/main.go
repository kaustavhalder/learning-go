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

// Anonymous Structs

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

	// Anonymous Structs 
	// no need for a placeholder
	// where to use this 
	
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "Vlad",
		last:  "rod",
		age:   33,
	}
	fmt.Println(p3)

}
