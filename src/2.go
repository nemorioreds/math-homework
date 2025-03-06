
// This program generates a random number between 1 and 10
package main

import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10) + 1
	fmt.Println("Your random number is: ", randomNumber)
}