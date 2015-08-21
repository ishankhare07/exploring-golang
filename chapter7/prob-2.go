package main

import "fmt"

func half(num int) (int, bool) {
    mid := num / 2
    if num % 2 == 0 {
        return mid, true
    } else {
        return mid, false
    }
}

func main() {
    var num int

    fmt.Print("Enter a number")
    fmt.Scanf("%d", &num)

    fmt.Println(half(num))
}
