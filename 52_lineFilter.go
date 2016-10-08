/*
Usage : cat writefile.txt | go run 52_lineFilter.go
Usage: go run 52_lineFilter.go : Then enter text via stdin
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		uppercaseStr := strings.ToUpper(str)
		fmt.Println(uppercaseStr)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
