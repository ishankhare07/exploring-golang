package main

import (
    "fmt"
    "math"
)

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (a *Rectangle) distance() float64 {
    l := a.x2 - a.x1
    h := a.y2 - a.y1
    return math.Sqrt(l*l + h*h)
}

func (a *Rectangle) area() float64 {
    l := a.x2 - a.x1
    h := a.y2 - a.y1

    return l*h
}

func main() {
    r := Rectangle{0,0,5,5}

    fmt.Println(r.distance())
    fmt.Println(r.area())
}
