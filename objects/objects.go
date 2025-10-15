package main

import (
	"fmt"
)

// An interface defines methods that the type must implement
type Animal interface {
	Speak()
	GetName() string
}

// Embedding Address struct to demonstrate composition
type Address struct {
	City    string
	Country string
}

type Cat struct {
	Address
	Name  string
	Color string
}

type Dog struct {
	Address
	Name  string
	Breed string
}

func (d Dog) Speak() {
	fmt.Printf("%s says: Woof!\n", d.Name)
}

func (d Dog) GetName() string {
	return d.Name
}

func (c Cat) Speak() {
	fmt.Printf("%s says: Meow!\n", c.Name)
}

func (c Cat) GetName() string {
	return c.Name
}