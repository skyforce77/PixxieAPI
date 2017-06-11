package main

import (
	"PixxieAPI"
	"time"
)

func OnInit() {

}

func OnEnable(bindings *PixxieAPI.Binding) {
	DrawTime(bindings)
}

func OnDisable(bindings *PixxieAPI.Binding) {

}

func OnTick(bindings *PixxieAPI.Binding) {
	DrawTime(bindings)
}

func DrawTime(bindings *PixxieAPI.Binding) {
	tim := time.Now()

	var i uint
	for i = 3; i >= 0 && i < 4; i-- {
		draw(bindings, i, tim.Hour(), tim.Minute(), tim.Second())
	}

	bindings.Push()
}

func draw(bindings *PixxieAPI.Binding, i uint, h int, m int, s int) {
	drawNb(bindings, i, 1, h)
	drawNb(bindings, i, 3, m)
	drawNb(bindings, i, 5, s)
}

func drawNb(bindings *PixxieAPI.Binding, y uint, x int, nb int) {
	n1 := nb/10
	n2 := nb%10

	if (n1 >> y) & 1 == 1 {
		bindings.SetPixel(x, int(y), 255, 255, 255)
	} else {
		bindings.SetPixel(x, int(y), 0, 0, 0)
	}
	if (n2 >> y) & 1 == 1 {
		bindings.SetPixel(x+1, int(y), 255, 255, 255)
	} else {
		bindings.SetPixel(x+1, int(y), 0, 0, 0)
	}
}

var (
	PluginDescriptor PixxieAPI.PixxiePlugin = PixxieAPI.PixxiePlugin{
		"binaryclock", 120,OnInit, OnEnable, OnDisable, OnTick,
	}
)