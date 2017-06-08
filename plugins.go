package PixxieAPI

import (
	"plugin"
	"io/ioutil"
	"fmt"
	"os"
	"container/list"
)

type PixxiePlugin struct {
	Id string
	OnInit func()
	OnEnable func(bindings *Binding)
	OnDisable func(bindings *Binding)
	OnTick func(bindings *Binding)
}

var (
	plugins list.List
)

func InitPlugins() {
	files, _ := ioutil.ReadDir("plugins")
	for _, f := range files {
		registerPlugin(f)
	}
}

func registerPlugin(f os.FileInfo) {
	plugin, err := plugin.Open(f.Name())
	check(err)

	descTest, err := plugin.Lookup("PluginDescriptor")
	check(err)

	desc, ok := descTest.(*PixxiePlugin)
	if !ok {
		panic(fmt.Sprintf("Can't convert PluginDescriptor from %s\n", f.Name()))
	}

	fmt.Printf("Loading plugin %s...\n", desc.Id)
	desc.OnInit()
	fmt.Printf("Successfuly loaded plugin %s.\n", desc.Id)
	plugins.PushBack(desc)
}
