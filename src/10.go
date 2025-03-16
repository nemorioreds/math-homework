package main

import (
	"fmt"
	"math/rand"
)

func randomGoCode() {
	// Generate a random number between 1 and 10
	randomNumber := rand.Intn(10) + 1

	// Print the generated number
	fmt.Println(randomNumber)
}
