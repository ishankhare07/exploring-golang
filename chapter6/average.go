package main

import "fmt"

func main() {
    var arr [5]float64
    var total float64 = 0

    for _, value := range arr {
        fmt.Scanf("%f", &value)
        total += value
    }

    fmt.Println("Average =", total/float64(len(arr)))
}
