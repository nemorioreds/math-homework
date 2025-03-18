package main

import "math/rand"

func randInt(min int, max int) int {
    return min + rand.Intn(max-min+1)
}

func main() {
	// Example usage:
	fmt.Println(randInt(0, 10)) // Output: 8
}
