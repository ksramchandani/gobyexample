/*
   Usage: go run 53_cliArgs.go abc def
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProgram := os.Args
	argsWithoutProgram := os.Args[1:]

	fmt.Println(argsWithProgram)
	fmt.Println(argsWithoutProgram)
}
