package main

import "fmt"
import "math"

type Circle struct {
    x, y, r float64
}

func areaCircle(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

func main() {
    var c Circle
    c.x = 0
    c.y = 0
    c.r = 10
    a := Circle{x: 0, y: 0, r: 5}

    fmt.Println(a.x, a.y, a.r)
    fmt.Println("Area = ", areaCircle(&a))
}
