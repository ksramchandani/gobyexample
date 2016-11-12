package main 

import (
    "fmt"
    "time"
)

func boring(msg string) <-chan string{
    c := make(chan string)

    go func(){
        for i:=0;; i++{
            c <- fmt.Sprintf("%s : %d", msg, i)
            time.Sleep(5*time.Second)
        }
    }()


    return c
}

func main() {

    joe := boring("joe")
    for i:=0; ;i++{
        select{
            case v1 := <-joe:
                fmt.Println(v1)
            case <- time.After(2*time.Second):
                fmt.Println("Timeout occured for wait")
                return
        }
    }
}
