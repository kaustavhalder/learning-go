package main

// to import a package, this needs to be inited with go.mod
// dependency management is done by go.sum

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speakAgent() {
	fmt.Println("this is from speak fucntion", s.first, s.last)
}

func (p person) speakPerson() {
	fmt.Println("this is from speak fucntion", p.first, p.last)
}

type humanInterface interface {
	speak()
}

// Anonymous functions
