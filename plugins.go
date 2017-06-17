package main

import (
	"plugin"
	"fmt"
	"os"
	"container/list"
	"time"
)

type PixxiePlugin struct {
	Id string
	DisplayTime int
	OnInit func()
	OnEnable func(bindings *Binding)
	OnDisable func(bindings *Binding)
	OnTick func(bindings *Binding)
}

var (
	def *PixxiePlugin
	active *PixxiePlugin
	plugins list.List
)

func InitPlugins() {
	/*files, _ := ioutil.ReadDir("plugins")
	for _, f := range files {
		registerPlugin(f)
	}*/

	PluginDescriptor.OnInit()
	plugins.PushBack(&PluginDescriptor)

	YoutubePluginDescriptor.OnInit()
	plugins.PushBack(&YoutubePluginDescriptor)

	TwitterPluginDescriptor.OnInit()
	plugins.PushBack(&TwitterPluginDescriptor)

	roll()

	cnt := 0
	for {
		time.Sleep(time.Millisecond*500)
		if active != nil {
			active.OnTick(bindings)

			if cnt % active.DisplayTime*2 == 0 {
				roll()
				cnt = 0
			}
		}
		cnt++;
	}
}

func registerPlugin(f os.FileInfo) {
	plugin, err := plugin.Open(fmt.Sprintf("plugins/%s", f.Name()))
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

func roll() {
	if active != nil {
		enabled := plugins.Front()
		plugins.Remove(enabled)
		plugins.PushBack(enabled.Value)
		disable(active)
	}
	enable(plugins.Front().Value.(*PixxiePlugin))
}

func enable(plugin *PixxiePlugin) {
	plugin.OnEnable(bindings)
	active = plugin
}

func disable(plugin *PixxiePlugin) {
	plugin.OnDisable(bindings)
	active = nil
	bindings.Clear()
}
