package main

import "fmt"

func main() {
    var temp float64

    fmt.Println("Enter temp in fahrenheit: ")
    fmt.Scanf("%f", &temp);

    fmt.Println((temp - 32) * 5/9);
}
