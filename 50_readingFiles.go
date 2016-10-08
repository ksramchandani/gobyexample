package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func separator() {
	fmt.Println("============================")
}

func main() {
	// Slurping the entire file contents into memory
	// Use this method only for very small files
	dataFile1, err := ioutil.ReadFile("readfile.txt")
	handleError(err)
	fmt.Println(string(dataFile1))
	separator()
	// Case that will cause a panic
	/*
		dataFile2, err := ioutil.ReadFile("junk.txt")
		handleError(err)
		fmt.Println(string(dataFile2))
	*/

	// Read a limited amount of bytes from the file
	f, err := os.Open("readfile.txt")
	b1 := make([]byte, 5)
	bytesRead, err := f.Read(b1)
	handleError(err)
	fmt.Println("Number of Bytes read = ", bytesRead)
	fmt.Println("Data of bytes read = ", string(b1))
	separator()

	/*
		Now you want to seek to a particular location in file
		and read from there
	*/
	loc, err := f.Seek(6, 0)
	handleError(err)
	b2 := make([]byte, 5)
	bytesRead, err = f.Read(b2)
	handleError(err)
	fmt.Println("Reading from location", loc)
	fmt.Println("Number of Bytes read = ", bytesRead)
	fmt.Println("Data of bytes read = ", string(b2))
	separator()

	/*
	   Robust implementation of data read from a file using io.ReadAtLeast
	*/

	loc2, err := f.Seek(6, 0) // Rewind i.e. This is needed to move the position back to where we want data to be read from
	handleError(err)
	b3 := make([]byte, 5)
	bytesRead, err = io.ReadAtLeast(f, b3, 5)
	fmt.Println("Reading from location", loc2)
	fmt.Println("Number of Bytes read = ", bytesRead)
	fmt.Println("Data of bytes read = ", string(b3))
	separator()

	_, err = f.Seek(0, 0)
	handleError(err)

	/*
	   Reading a file using bufio
	*/

	reader := bufio.NewReader(f)
	b4 := make([]byte, 5)
	b4, err = reader.Peek(5)
	fmt.Println("Data of bytes read = ", string(b4))

	f.Close()
}
