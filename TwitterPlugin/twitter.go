package main

import (
	"github.com/skyforce77/PixxieAPI"
	"log"
	"github.com/PuerkitoBio/goquery"
)

var (
	c1 = PixxieAPI.Pixel{0x55, 0x84, 0xba}
	c2 = PixxieAPI.Pixel{0xc9, 0xca, 0x61}
	c3 = PixxieAPI.Pixel{0x00, 0x00, 0x00}

	logo = []PixxieAPI.Pixel{
		c3, c3, c3, c1,
		c1, c3, c1, c2,
		c1, c1, c1, c1,
		c3, c1, c1, c3,
	}

	Id = "twitter"

	cnt = 0
	followers = "0"

	refreshTime = 0
	refreshCnt = 0
)

func OnInit() {
	refreshTime = int(PixxieAPI.GetPluginConfig(Id, "refresh").(float64))
	refreshTime *= 120
	refresh()
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

	refreshCnt++
	if refreshCnt == refreshTime {
		refresh()
		refreshCnt = 0;
	}
}

func refresh() {
	doc, err := goquery.NewDocument(PixxieAPI.GetPluginConfig(Id, "user").(string))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".ProfileNav-item--followers > .ProfileNav-stat > .ProfileNav-value").Each(func(i int, s *goquery.Selection) {
		content, err := s.Html()
		if err == nil {
			followers = content
		}
	})
}

var (
	PluginDescriptor PixxieAPI.PixxiePlugin = PixxieAPI.PixxiePlugin{
		Id, 60, OnInit, OnEnable, OnDisable, OnTick,
	}
)
