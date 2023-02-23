package main

// to import a package, this needs to be inited with go.mod
// dependency management is done by go.sum

import "fmt"

type personA struct {
	first string
	last  string
}

type secretAgentA struct {
	personA
	ltk bool
}

func (s secretAgentA) speakAgent() {
	fmt.Println("this is from speak fucntion", s.first, s.last)
}

func (p personA) speakPerson() {
	fmt.Println("this is from speak fucntion", p.first, p.last)
}

type humanInterface interface {
	speak()
}

// Anonymous functions
