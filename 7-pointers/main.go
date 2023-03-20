package main

// This is the part which shows pointers
import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	// Every type has its own pointer (type)
	// everything in go is shared by value
	// & gives us the memory address
	// * dereferences and gives us the value from the pointer
	fmt.Println("This section is for pointers")
	a := 42
	fmt.Println(a)
	// Below gives us the memory location of the variable a
	fmt.Printf("%T\n", a)
	// This is a pointer to the memory location
	fmt.Printf("%T\n", &a)
	fmt.Println("Dereferencing")
	fmt.Printf("%T\n", *&a)
	name := "Vlad"
	fmt.Println(name)
	// Below gives us the memory location of the variable a
	fmt.Println(&name)
	fmt.Printf("%T\n", name)
	// This is a pointer to the memory location
	fmt.Printf("%T\n", &name)
	// Method sets
	// Types can have methods attached to them
}
