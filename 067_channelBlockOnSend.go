package main 

import (
    "fmt"
)

func main() {
    c := make(chan bool)

    // Send to channel
    /* This will be a blocking call 
    until a goroutine receives from this channel */
    c <- true

    fmt.Println("This will never print")
}
