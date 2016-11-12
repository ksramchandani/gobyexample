package main

import (
	"fmt"
)

const (
	WHITE = iota
	RED
	BLUE
	GREEN
)

const (
	one = 1 << iota
	two
	four
	eight
	sixteen
)

const (
	mute = iota
	mono
	stereo
	_
	_
	surround
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // Kilobyte
	MB
	GB
	TB
	PB
	EB // Exabyte
	ZB // Zettayte
	YB //Yottabyte
)

func separator() {
	fmt.Println("=============")
}

func main() {
	fmt.Println("WHITE", WHITE)
	fmt.Println("RED", RED)
	fmt.Println("BLUE", BLUE)
	fmt.Println("GREEN", GREEN)
	separator()

	fmt.Println("one", one)
	fmt.Println("two", two)
	fmt.Println("four", four)
	fmt.Println("eight", eight)
	fmt.Println("sixteen", sixteen)
	separator()

	fmt.Println("mute", mute)
	fmt.Println("mono", mono)
	fmt.Println("stereo", stereo)
	fmt.Println("surround", surround)
	separator()

	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)
	fmt.Println("TB", TB)
	fmt.Println("PB", PB)
	fmt.Println("EB", EB)
	fmt.Println("ZB", ZB)
	fmt.Println("YB", YB)

}
