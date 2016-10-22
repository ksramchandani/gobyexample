package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// First find path of binary

	lsPath, _ := exec.LookPath("ls")
	fmt.Println(lsPath)

	args := []string{"ls", "-l", "-r", "-t", "-h"}

	env := os.Environ()

	if err := syscall.Exec(lsPath, args, env); err != nil {
		panic(err)
	}

}
