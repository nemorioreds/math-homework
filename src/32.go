package main

import (
    "fmt"
    "time"
)

func main() {
    var score int = 100
    start := time.Now()
    fmt.Printf("Your current score: %d\n", score)
    end := time.Now()

    // Add a small delay to simulate a time-consuming task
    time.Sleep(5 * time.Second)

    // Calculate the elapsed time in seconds since the beginning of the program
    duration := end.Sub(start)

    // Format and print the total elapsed time
    fmt.Printf("Total elapsed time: %s\n", duration)
}
