package PixxieAPI

import (
	"plugin"
)

type Binding struct {
	SetPixel func(x int, y int, r int, g int, b int)
	Brightness func(bright float32)
}

var (
	bindings *Binding
)

func InitBindings() {
	plugin, err := plugin.Open("bindings/unicorn.so")
	check(err)

	bindTest, err := plugin.Lookup("Bindings")
	check(err)

	bind, ok := bindTest.(*Binding)
	if !ok {
		panic("Can't convert Bindings")
	}
	bindings = bind
}
