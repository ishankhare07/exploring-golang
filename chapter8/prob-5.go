package main

import "fmt"

func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    var a,b int
    fmt.Println("Enter two numbers")
    fmt.Scanf("%d %d", &a, &b)
    fmt.Println("a = ", a, " b = ", b)
    swap(&a, &b)
    fmt.Println("a = ", a, " b = ", b)
}
