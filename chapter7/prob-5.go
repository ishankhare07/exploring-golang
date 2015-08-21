//fibonacci sqeuence recursive

package main

import "fmt"

func fib(num int) int {
    if num == 0 {
        return 0
    } else if num == 1 {
        return 1
    } else {
        return fib(num - 1) + fib(num - 2)
    }
}

func main() {
    var num int
    fmt.Print("enter a number: ")
    fmt.Scanf("%d", &num)

    fmt.Println(fib(num))
}
