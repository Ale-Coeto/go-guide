package main

import (
	"fmt"
)

func main() {
	// Create a cat instance
	cat := Cat{
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
		Name:  "Whiskers",
		Color: "Black",
	}
	
	// Create a dog instance
	dog := Dog{
		Address: Address{
			City:    "Los Angeles",
			Country: "USA",
		},
		Name:  "Buddy",
		Breed: "Golden Retriever",
	}

	// Use them as Animal interface (polymorphism)
	animalList := []Animal{cat, dog}

	fmt.Println("=== Animals speaking ===")
	for _, animal := range animalList {
		fmt.Printf("Animal Name: %s\n", animal.GetName())
		animal.Speak()
		fmt.Println()
	}

	// Composition example
	fmt.Println("=== Animal Locations ===")
	fmt.Printf("%s lives in %s, %s\n", cat.GetName(), cat.City, cat.Country) // Composition access
	fmt.Printf("%s lives in %s, %s\n", dog.GetName(), dog.Address.City, dog.Address.Country)

}