package main

import (
    "fmt"
    "math"
)

type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {       //in between func and name of function, we have added a receiver
    return math.Pi * c.r * c.r
}

func main() {
    c := Circle{0, 0, 10}
    fmt.Println("Area = ", c.area())
}
