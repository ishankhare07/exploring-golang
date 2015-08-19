package main

import "os"
import "fmt"

func main() {
    fmt.Println("exiting now...")
    os.Exit(0)
    fmt.Println("this should not print")
}
