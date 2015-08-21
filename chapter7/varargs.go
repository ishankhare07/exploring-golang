package main

import "fmt"

func add(args ...int) int {
    total := 0
    for _, value := range args {
        total += value
    }

    return total
}

func addArrayAsArgs(args []int) int {
    total := 0
    for _, value := range args {
        total += value
    }

    return total
}

func main() {
    list := []int{1,2,3,4,5}
    fmt.Println(add(list...))                   //argument unpacking
    fmt.Println(addArrayAsArgs(list))           //without unpacking, send array as argument
}
