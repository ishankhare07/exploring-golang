package main

import "fmt"

func main() {
    var feet float64

    fmt.Println("Enter height in feet: ")
    fmt.Scanf("%f", &feet)

    fmt.Println(feet * 0.3048)
}
