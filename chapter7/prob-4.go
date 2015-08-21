package main

import "fmt"

func makeOddGenerator() func() int {
    current := 1

    return func() (ret int) {
        ret = current
        current += 2

        return
    }
}

func main() {
    var limit int

    nextOdd := makeOddGenerator()
    fmt.Print("Enter the upper limit: ")
    fmt.Scanf("%d", &limit)

    for i := 0; i <= limit/2; i++ {
        fmt.Println(nextOdd())
    }
}
