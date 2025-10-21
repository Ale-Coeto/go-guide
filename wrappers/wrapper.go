package main

import (
	"fmt"
)

func wrapper(f func (string)) {
	utilityString := "Important data"

	f(utilityString)

}

func embeddingWrapper() {
	wrapper(func(info string) {
		fmt.Println("Using wrapper:", info)
	})
}

func innerFunction(info string) {
	fmt.Println("Inner function received data: ", info)
}

func main() {
	wrapper(innerFunction)
	embeddingWrapper()
}