package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	/* Write string as bytes to file*/
	f, err := os.Create("writefile.txt")
	bytesWritten, err := f.Write([]byte("WriteLine1\n"))
	handleError(err)
	fmt.Println("Number of bytesWritten = ", bytesWritten)
	f.Sync()

	/* Write string to file*/
	bytesWritten, err = f.WriteString("WriteLine2\n")
	handleError(err)
	fmt.Println("Number of bytesWritten = ", bytesWritten)
	f.Sync()

	/* Write using bufio */
	w := bufio.NewWriter(f)
	bytesWritten, err = w.WriteString("WriteLine3\n")
	handleError(err)
	fmt.Println("Number of bytesWritten = ", bytesWritten)
	w.Flush()

	defer f.Close()
}
