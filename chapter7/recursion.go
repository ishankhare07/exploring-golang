package main

import "fmt"

func factorial(fact int) int {
    if fact <= 0 {
        return 1
    } else {
        return fact * factorial(fact - 1)
    }
}

func main() {
    var fact int

    fmt.Print("Enter a number: ")
    fmt.Scanf("%d", &fact)

    fmt.Println(factorial(fact))
}
