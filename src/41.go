package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter an integer: ")
    fmt.Scan(&n)

    if n%2 == 0 {
        fmt.Println("The number is even.")
    } else {
        fmt.Println("The number is odd.")
    }
}
