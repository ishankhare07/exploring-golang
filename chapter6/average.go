package main

import "fmt"

func main() {
    var arr [5]float64
    var total float64 = 0

    for i := 0; i < len(arr); i++ {
        fmt.Scanf("%f", &arr[i])
        total += arr[i]
    }

    fmt.Println("Average =", total/float64(len(arr)))
}
