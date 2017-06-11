package PixxieAPI

import "fmt"

var (
	letters = make(map[rune][]Pixel)
	b = Pixel{0,0,0}
	w = Pixel{255, 255, 255}
)

func InitString() {
	letters[' '] = []Pixel{
		b, b, b, b,
		b, b, b, b,
		b, b, b, b,
		b, b, b, b,
	}

	letters['0'] = []Pixel{
		w, w, w, w,
		w, w, b, w,
		w, b, w, w,
		w, w, w, w,
	}
	letters['1'] = []Pixel{
		w, w, w, b,
		b, w, w, b,
		b, w, w, b,
		w, w, w, w,
	}
	letters['2'] = []Pixel{
		w, w, w, w,
		b, w, w, w,
		b, w, w, b,
		w, w, w, w,
	}
	letters['3'] = []Pixel{
		w, w, w, w,
		b, w, w, w,
		b, b, w, w,
		w, w, w, w,
	}
	letters['4'] = []Pixel{
		w, b, w, w,
		w, w, w, w,
		w, w, w, w,
		b, b, w, w,
	}
	letters['5'] = []Pixel{
		w, w, w, w,
		w, w, w, b,
		b, w, w, w,
		w, w, w, w,
	}
	letters['6'] = []Pixel{
		w, b, b, b,
		w, w, w, w,
		w, b, b, w,
		w, w, w, w,
	}
	letters['7'] = []Pixel{
		w, w, w, w,
		b, b, w, w,
		b, b, w, w,
		b, b, w, w,
	}
	letters['8'] = []Pixel{
		b, w, w, w,
		b, w, w, w,
		w, b, w, w,
		w, w, w, w,
	}
	letters['9'] = []Pixel{
		w, w, w, w,
		w, b, w, w,
		w, w, w, w,
		b, b, w, w,
	}
}

func (bindings *Binding) DrawString(x int, str string) {
	for i, c := range str {
		if template, ok := letters[c]; ok {
			bindings.Draw(i*5-x, 0, template, 4, 4)
		}
	}
}

func DisplayedSize(str string) int {
	return len(str)*5
}