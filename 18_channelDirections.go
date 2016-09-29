package main

import (
	"fmt"
)

func sendToPings(pings chan<- string, msg string) {
	fmt.Printf("In sendToPings.. Sending %s to pings\n", msg)
	pings <- msg
}

func sendFromPingsToPongs(pings <-chan string, pongs chan<- string) {
	fmt.Println("In sendFromPingsToPongs.. Sending from pings to pongs")
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	sendToPings(pings, "abcdef")
	sendFromPingsToPongs(pings, pongs)
	fmt.Println("pongs got", <-pongs)

}
