package main

import "fmt"

func main() {
	fmt.Println("=== Using Real Person ===")
	// Create real instance
	realPerson := RealPerson{Name: "Alice", Age: 30}
	service := SomeService{
		person: realPerson,
	}
	service.Run()

	fmt.Println("\n=== Using Mock Person ===")
	// Create mock instance (for testing)
	mockPerson := MockPerson{Name: "Bob", Age: 25}
	mockService := SomeService{
		person: mockPerson,
	}
	mockService.Run()
}