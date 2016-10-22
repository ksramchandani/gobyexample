package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Simple command which takes no arguments
	dateCmd := exec.Command("date")
	dateStr, err := dateCmd.Output()
	handleError(err)
	fmt.Println(string(dateStr))

	// Piping data to external processes
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\nno grep"))
	grepIn.Close()

	bytesRead, err := ioutil.ReadAll(grepOut)
	fmt.Println(string(bytesRead))

	// Providing arguments as a single string
	lsCmd := exec.Command("bash", "-c", "ls -lrth")
	lsOutput, err := lsCmd.Output()
	handleError(err)
	fmt.Println(string(lsOutput))

}
