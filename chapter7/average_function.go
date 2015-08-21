package main

import "fmt"

func average(xs []float64) float64 {
    var total float64 = 0.0

    for _, value := range xs {
        total += value
    }

    return total / float64(len(xs))
}

func main() {
    xs := []float64{23,77,98,45,95}

    fmt.Println(average(xs))
}
