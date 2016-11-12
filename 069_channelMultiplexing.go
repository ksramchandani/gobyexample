/* From https://youtu.be/f6kdp27TYZs?t=965 */

package main 

import (
    "fmt"
    "time"
)

func boring(msg string) <-chan string {
    c := make(chan string)

    go func() {
        for i:=0; ;i++ {
            c <- fmt.Sprintf("%s : %d", msg, i)
            if msg == "joe" {
                time.Sleep(time.Second)
            }
            time.Sleep(time.Second)
        }
    }()
    return c
}

func multiplexor(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func(){
        for {
            c <- <-input1
        }
    }()

    go func(){
        for {
            c <- <-input2
        }
    }()

    return c
}

func main() {
    joe := boring("joe")
    ann := boring("ann")
    c := multiplexor(joe, ann)

    for i:=0; i< 10; i++ {
        fmt.Println(<-c)
    }
}
