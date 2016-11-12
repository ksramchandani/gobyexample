package main

import (
	"fmt"
)

const (
	WHITE = iota // first element
	BLACK
	BLU
	REDD
	YELLOW
)

type Color byte // Color value

type Box struct {
	width  float64
	height float64
	depth  float64
	color  Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (c Color) String() string {
	values := []string{"White", "Black", "Blue", "Red", "Yellow"}
	return values[c]
}

func (bl BoxList) BiggestBox() {
	var biggestVolume float64
	var biggestBoxColor string
	for _, box := range bl {
		v := box.Volume()
		if v > biggestVolume {
			biggestVolume = v
			biggestBoxColor = box.color.String()
		}
	}
	fmt.Println(biggestVolume, biggestBoxColor)
}

func (bPtr *Box) PaintBoxBlack() {
	bPtr.color = 1
}

func (bl BoxList) PaintBoxListBlack() {
	for i, _ := range bl {
		bl[i].PaintBoxBlack()
		fmt.Printf("Painted box with volume %f %s\n", bl[i].Volume(), bl[i].color.String())
	}
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, REDD},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLU},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Println("total boxes", len(boxes))
	fmt.Println("volume of first one", boxes[0].Volume())
	fmt.Println("color of last one is", boxes[len(boxes)-1].color.String())
	boxes.BiggestBox()
	boxes.PaintBoxListBlack()
	boxes.BiggestBox()
}
