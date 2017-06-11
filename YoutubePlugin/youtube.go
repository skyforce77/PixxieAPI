package main

import (
	"PixxieAPI"
)

var (
	c1 = PixxieAPI.Pixel{0xa9, 0x1e, 0x1e}
	c2 = PixxieAPI.Pixel{0xff, 0xff, 0xff}
	c3 = PixxieAPI.Pixel{0xc5, 0x68, 0x68}
	c4 = PixxieAPI.Pixel{0xcc, 0x92, 0x92}

	logo = []PixxieAPI.Pixel{
		c1, c1, c1, c1,
		c1, c2, c3, c1,
		c1, c2, c4, c1,
		c1, c1, c1, c1,
	}
)

func OnInit() {

}

func OnEnable(bindings *PixxieAPI.Binding) {
	bindings.Draw(3, 0, logo, 4, 4)
	bindings.Push()
}

func OnDisable(bindings *PixxieAPI.Binding) {

}

func OnTick(bindings *PixxieAPI.Binding) {

}

var (
	PluginDescriptor PixxieAPI.PixxiePlugin = PixxieAPI.PixxiePlugin{
		"youtube", 30, OnInit, OnEnable, OnDisable, OnTick,
	}
)
