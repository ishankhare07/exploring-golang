package main

import "fmt"

func main() {
    x := new(string)
    fmt.Print("Enter a string: ")
    fmt.Scanf("%s", x)
    fmt.Println(*x)
}
