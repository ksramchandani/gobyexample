package main 

import (
    "fmt"
)

func main() {
    c := make(chan bool)

    // Receive from channel
    // This is a blocking call
    <-c

    fmt.Println("This will never print")
}
