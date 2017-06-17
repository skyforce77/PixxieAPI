package main

var (
	letters = make(map[rune][]Pixel)
	b = Pixel{0,0,0}
	w = Pixel{255, 255, 255}
)

func (bindings *Binding) DrawString(x int, str string) {
	for i, c := range str {
		if i*5-x <= -4 {
			continue
		}
		if i*5-x >= 8 {
			return
		}
		if template, ok := letters[c]; ok {
			bindings.Draw(i*5-x, 0, template, 4, 4)
		}
	}
}

func DisplayedSize(str string) int {
	return len(str)*5
}

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
		w, w, w, b,
		w, b, w, w,
		w, w, b, w,
		b, w, w, w,
	}
	letters['9'] = []Pixel{
		w, w, w, w,
		w, b, w, w,
		w, w, w, w,
		b, b, w, w,
	}

	letters['A'] = []Pixel{
		w, w, w, w,
		w, w, b, w,
		w, w, w, w,
		w, w, b, w,
	}
	letters['B'] = []Pixel{
		w, w, w, b,
		w, b, w, w,
		w, w, b, w,
		w, w, w, w,
	}
	letters['C'] = []Pixel{
		w, w, w, w,
		w, w, b, b,
		w, w, b, b,
		w, w, w, w,
	}
	letters['D'] = []Pixel{
		w, w, w, b,
		w, w, b, w,
		w, w, b, w,
		w, w, w, b,
	}
	letters['E'] = []Pixel{
		w, w, w, w,
		w, w, w, b,
		w, w, b, b,
		w, w, w, w,
	}
	letters['F'] = []Pixel{
		w, w, w, w,
		w, w, w, b,
		w, w, w, w,
		w, w, w, b,
	}
	letters['G'] = []Pixel{
		w, w, w, w,
		w, w, b, b,
		w, w, b, w,
		w, w, w, w,
	}
	letters['H'] = []Pixel{
		w, w, b, w,
		w, w, w, w,
		w, w, w, w,
		w, w, b, w,
	}
	letters['I'] = []Pixel{
		w, w, w, w,
		b, w, w, b,
		b, w, w, b,
		w, w, w, w,
	}
	letters['J'] = []Pixel{
		b, b, w, w,
		b, b, w, w,
		w, w, w, w,
		w, w, w, w,
	}
	letters['K'] = []Pixel{
		w, w, b, w,
		w, w, w, b,
		w, w, w, w,
		w, w, b, w,
	}
	letters['L'] = []Pixel{
		w, w, w, b,
		w, w, w, b,
		w, w, w, b,
		w, w, w, w,
	}
	letters['M'] = []Pixel{
		w, w, w, w,
		w, b, w, w,
		w, b, b, w,
		w, b, b, w,
	}
	letters['N'] = []Pixel{
		w, w, b, w,
		w, w, w, w,
		w, w, w, w,
		w, b, w, w,
	}
	letters['O'] = []Pixel{
		w, w, w, w,
		w, w, b, w,
		w, w, b, w,
		w, w, w, w,
	}
	letters['P'] = []Pixel{
		w, w, w, w,
		w, b, b, w,
		w, w, w, w,
		w, w, b, b,
	}
	letters['Q'] = []Pixel{
		w, w, w, w,
		w, b, b, w,
		w, b, w, w,
		w, w, w, w,
	}
	letters['R'] = []Pixel{
		w, w, w, w,
		w, w, b, w,
		w, w, w, b,
		w, w, b, w,
	}
	letters['S'] = []Pixel{
		w, w, w, w,
		w, w, b, b,
		b, b, w, w,
		w, w, w, w,
	}
	letters['T'] = []Pixel{
		w, w, w, w,
		w, w, w, w,
		b, w, w, b,
		b, w, w, b,
	}
	letters['U'] = []Pixel{
		w, b, b, w,
		w, b, b, w,
		w, w, w, w,
		w, w, w, w,
	}
	letters['V'] = []Pixel{
		w, b, b, w,
		w, b, b, w,
		w, w, w, w,
		b, w, w, b,
	}
	letters['W'] = []Pixel{
		w, b, b, w,
		w, w, b, w,
		w, w, w, w,
		w, w, w, w,
	}
	letters['X'] = []Pixel{
		w, b, b, w,
		w, b, b, w,
		b, w, w, b,
		w, b, b, w,
	}
	letters['Y'] = []Pixel{
		w, b, b, w,
		w, b, b, w,
		b, w, w, b,
		b, w, w, b,
	}
	letters['Z'] = []Pixel{
		w, w, w, w,
		b, b, w, w,
		w, w, b, b,
		w, w, w, w,
	}

	letters['.'] = []Pixel{
		b, b, b, b,
		b, b, b, b,
		b, b, b, b,
		b, w, w, b,
	}
	letters[','] = []Pixel{
		b, b, b, b,
		b, b, b, b,
		b, b, w, b,
		b, w, w, b,
	}
	letters['['] = []Pixel{
		b, w, w, b,
		b, w, b, b,
		b, w, b, b,
		b, w, w, b,
	}
	letters[']'] = []Pixel{
		b, w, w, b,
		b, b, w, b,
		b, b, w, b,
		b, w, w, b,
	}
	letters['('] = []Pixel{
		b, b, w, b,
		b, w, b, b,
		b, w, b, b,
		b, b, w, b,
	}
	letters[')'] = []Pixel{
		b, w, b, b,
		b, b, w, b,
		b, b, w, b,
		b, w, b, b,
	}
	letters['^'] = []Pixel{
		b, w, b, b,
		w, b, w, b,
		b, b, b, b,
		b, b, b, b,
	}
	letters['<'] = []Pixel{
		b, b, w, b,
		b, w, b, b,
		b, b, w, b,
		b, b, b, b,
	}
	letters['>'] = []Pixel{
		b, w, b, b,
		b, b, w, b,
		b, w, b, b,
		b, b, b, b,
	}
	letters['!'] = []Pixel{
		b, w, w, b,
		b, w, w, b,
		b, b, b, b,
		b, w, w, b,
	}
	letters['?'] = []Pixel{
		w, w, w, b,
		w, b, w, b,
		b, b, b, b,
		b, w, w, b,
	}
	letters['"'] = []Pixel{
		w, b, w, b,
		w, b, w, b,
		b, b, b, b,
		b, b, b, b,
	}
	letters['\''] = []Pixel{
		b, w, w, b,
		b, b, w, b,
		b, b, b, b,
		b, b, b, b,
	}
	letters[';'] = []Pixel{
		b, w, b, b,
		b, b, b, b,
		b, w, b, b,
		b, w, b, b,
	}
	letters[':'] = []Pixel{
		b, w, b, b,
		b, b, b, b,
		b, w, b, b,
		b, b, b, b,
	}
	letters['~'] = []Pixel{
		b, b, b, b,
		b, w, b, w,
		w, b, w, b,
		b, b, b, b,
	}
	letters['&'] = []Pixel{
		b, b, w, b,
		w, w, w, w,
		w, b, w, b,
		w, w, w, b,
	}
	letters['@'] = []Pixel{
		w, w, w, w,
		w, b, w, w,
		w, b, b, b,
		w, w, w, w,
	}
	letters['*'] = []Pixel{
		b, w, b, b,
		w, b, w, b,
		b, w, b, b,
		b, b, b, b,
	}
	letters['|'] = []Pixel{
		b, w, b, b,
		b, w, b, b,
		b, w, b, b,
		b, w, b, b,
	}
	letters['_'] = []Pixel{
		b, b, b, b,
		b, b, b, b,
		b, b, b, b,
		w, w, w, w,
	}
	letters['\\'] = []Pixel{
		w, b, b, b,
		b, w, b, b,
		b, b, w, b,
		b, b, b, w,
	}
	letters['%'] = []Pixel{
		w, b, b, w,
		b, b, w, b,
		b, w, b, b,
		w, b, b, w,
	}
	letters['+'] = []Pixel{
		b, w, w, b,
		w, w, w, w,
		b, w, w, b,
		b, b, b, b,
	}
	letters['-'] = []Pixel{
		b, b, b, b,
		b, w, w, b,
		b, b, b, b,
		b, b, b, b,
	}
	letters['/'] = []Pixel{
		b, b, b, w,
		b, b, w, b,
		b, w, b, b,
		w, b, b, b,
	}
	letters['='] = []Pixel{
		b, b, b, b,
		w, w, w, w,
		b, b, b, b,
		w, w, w, w,
	}
}