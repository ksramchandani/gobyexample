package main 

import (
    "fmt"
    "time"
)

func boring(msg string) <-chan string {
    c := make(chan string)

    go func() {
        for i:=0; ; i++ {
            c <- fmt.Sprintf("%s : %d", msg, i)
            if msg == "joe" {
                time.Sleep(time.Second)            
            }
            time.Sleep(time.Second)
        }
    }()
    return c
}

func main() {

    joe := boring("joe")
    ann := boring("ann")
    for i:=0; i < 5; i++ {
        fmt.Println(<-joe)
        fmt.Println(<-ann)
    }

}
