package PixxieAPI

import (
	"plugin"
	"fmt"
)

type Pixel struct {
	R, G, B int
}

type Binding struct {
	SetPixel func(x int, y int, r int, g int, b int)
	Brightness func(bright float32)
	Push func()
	Draw func(x int, y int, image []Pixel, width int, height int)
	Clear func()
}

var (
	bindings *Binding
)

func InitBindings() {
	plugin, err := plugin.Open(fmt.Sprintf("bindings/%s", config["bindings"]))
	check(err)

	bindTest, err := plugin.Lookup("Bindings")
	check(err)

	bind, ok := bindTest.(*Binding)
	if !ok {
		panic("Can't convert Bindings")
	}
	bindings = bind
}
