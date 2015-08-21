package main

import "fmt"
import "strconv"

func greatest(args ...int) (greatest int) {
    greatest = args[0]
    for _, value := range args {
        if greatest < value {
            greatest = value
        }
    }

    return
}

func main() {
    var list [10]int
    for i := 0; i < 10; i++ {
        fmt.Print("Enter element #" + strconv.Itoa(i))
        fmt.Scanf("%d", &list[i])
    }

    fmt.Println(greatest(list[:]...))
}
