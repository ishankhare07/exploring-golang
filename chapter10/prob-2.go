/*sleep function implementatino using time.After*/

package main

import (
    "fmt"
    "time"
)

func mySleep(duration int64) {
    select {
        case <- time.After(time.Second * time.Duration(duration)):
            return
    }
}

func main() {
    var time int64
    fmt.Print("Enter sleep duration: ")
    fmt.Scanf("%d", &time)
    mySleep(time)
    fmt.Println("awake")
}
