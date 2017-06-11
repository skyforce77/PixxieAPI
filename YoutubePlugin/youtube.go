package main

import (
	"PixxieAPI"
	"log"
	"github.com/PuerkitoBio/goquery"
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

	cnt = 0
	followers = "0"
)

func OnInit() {
	doc, err := goquery.NewDocument(PixxieAPI.GetPluginConfig("youtube", "channel").(string))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".yt-subscription-button-subscriber-count-branded-horizontal").Each(func(i int, s *goquery.Selection) {
		content, err := s.Html()
		if err == nil {
			followers = content
		}
	})
}

func OnEnable(bindings *PixxieAPI.Binding) {

}

func OnDisable(bindings *PixxieAPI.Binding) {

}

func OnTick(bindings *PixxieAPI.Binding) {
	bindings.Clear()
	bindings.Draw(-cnt, 0, logo, 4, 4)
	bindings.DrawString(cnt-5, followers)
	bindings.Push()

	cnt++
	if cnt > 5+PixxieAPI.DisplayedSize(followers) {
		cnt = 0
	}
}

var (
	PluginDescriptor PixxieAPI.PixxiePlugin = PixxieAPI.PixxiePlugin{
		"youtube", 60, OnInit, OnEnable, OnDisable, OnTick,
	}
)
