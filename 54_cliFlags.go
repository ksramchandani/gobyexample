/*
   Usage:
        go run 54_cliFlags.go
        go run 54_cliFlags.go -sflag=def -iflag=25 -bflag -vflag=xxx
        go run 54_cliFlags.go -sflag=def -iflag=25
        go run 54_cliFlags.go -sflag=def
        go run 54_cliFlags.go -sflag=def -iflag=25 -vflag=xxx x y z
        go run 54_cliFlags.go -sflag=def -iflag=25 x y z -bflag
        go run 54_cliFlags.go -sflag=def -iflag=25 -bflag x y z
        go run 54_cliFlags.go -jflag=10

*/

package main

import (
	"flag"
	"fmt"
)

func main() {
	stringPtr := flag.String("sflag", "abc", "String Flag")
	intPtr := flag.Int("iflag", 10, "Int Flag")
	boolPtr := flag.Bool("bflag", false, "Bool Flag")

	var vflag string
	flag.StringVar(&vflag, "vflag", "vvv", "String Flag")

	flag.Parse()

	fmt.Println("sflag:", *stringPtr)
	fmt.Println("iflag:", *intPtr)
	fmt.Println("bflag:", *boolPtr)
	fmt.Println("vflag:", vflag)
	fmt.Println("tail:", flag.Args())

}
