package main

import (
    "fmt"
    "time"
    "math/rand"
)

func f(x int) {
    for i := 0; i < 10; i++ {
        fmt.Println(x, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    for x := 0; x < 10; x++ {
        go f(x)
    }
    var input string
    fmt.Scanln(&input)
}
