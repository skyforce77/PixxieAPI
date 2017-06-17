package main

import (
	"fmt"
)

func SetPixel(x int, y int, r int, g int, b int) {
	index := x+8*y
	if x < 0 || x > 7 || y < 0 || y > 3 {
		return;
	}

	if r+g+b == 0 {
		data[index] = false
	} else {
		data[index] = true
	}
}

func Draw(x int, y int, image []Pixel, width int, height int) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			p := image[i+width*j]
			SetPixel(i+x, j+y, p.R, p.G, p.B)
		}
	}
}

func Brightness (bright float32) {

}

func Push() {
	fmt.Print("\033c")
	for i := 0; i < 4*8; i++ {
		if data[i] {
			fmt.Print("██")
		} else {
			fmt.Print("  ")
		}

		if i%8 == 7 {
			fmt.Println()
		}
	}
}

func Clear() {
	for i := 0; i < 4*8; i++ {
		data[i] = false
	}
}

var (
	Bindings Binding = Binding{
		SetPixel,
		Brightness,
		Push,
		Draw,
		Clear,
	}
	data []bool = make([]bool, 4*8)
)
