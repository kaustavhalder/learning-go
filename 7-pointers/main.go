package main

import "fmt"

func main() {
	// Every type has its own pointer (type)
	fmt.Println("This section is for pointers")
	a := 42
	fmt.Println(a)
	// Below gives us the memory location of the variable a
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	// This is a pointer to the memory location
	fmt.Printf("%T\n", &a)

	name := "Vlad"
	fmt.Println(name)
	// Below gives us the memory location of the variable a
	fmt.Println(&name)
	fmt.Printf("%T\n", name)
	// This is a pointer to the memory location
	fmt.Printf("%T\n", &name)
}
