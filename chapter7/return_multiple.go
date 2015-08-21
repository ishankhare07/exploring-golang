package main

import "fmt"

func getNumbers() (int, int) {
    return 5,6
}

func main() {
    x, y := getNumbers()
    fmt.Println(x,y)
}
