package main

import (
    "fmt"
    "time"
)

func pinger(c chan<- string) {             // can only send to channel
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func ponger(c chan <- string) {             // can only send to channel
    for {
        c <- "pong"
    }
}

func printer(c <-chan string) {             // can only receive from channel
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)

    var input string
    fmt.Println("Press enter to exit")
    fmt.Scanln(&input)
}
