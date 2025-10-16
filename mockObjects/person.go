package main

import (
	"fmt"
)

// Define an interface (like azureclients.Queue in your working example)
type Person interface {
	Greet()
}

type RealPerson struct {
	Name string
	Age  int
}

func (r RealPerson) Greet() {
	fmt.Printf("Hello, I am %s and I am %d years old\n", r.Name, r.Age)
}

type MockPerson struct {
	Name string
	Age  int
}

func (m MockPerson) Greet() {
	fmt.Printf("Hello, I am %s and I am %d years old (from MockPerson)\n", m.Name, m.Age)
}

type SomeService struct {
	person Person
}

func (s SomeService) Run() {
	s.person.Greet()
}
