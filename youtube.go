package main

import (
	"log"
	"github.com/PuerkitoBio/goquery"
)

var (
	Youtubec1 = Pixel{0xa9, 0x1e, 0x1e}
	Youtubec2 = Pixel{0xff, 0xff, 0xff}
	Youtubec3 = Pixel{0xc5, 0x68, 0x68}
	Youtubec4 = Pixel{0xcc, 0x92, 0x92}

	Youtubelogo = []Pixel{
		Youtubec1, Youtubec1, Youtubec1, Youtubec1,
		Youtubec1, Youtubec2, Youtubec3, Youtubec1,
		Youtubec1, Youtubec2, Youtubec4, Youtubec1,
		Youtubec1, Youtubec1, Youtubec1, Youtubec1,
	}

	YoutubeId = "youtube"

	Youtubecnt = 0
	Youtubefollowers = "0"

	YoutuberefreshTime = 0
	YoutuberefreshCnt = 0
)

func YoutubeOnInit() {
	YoutuberefreshTime = int(GetPluginConfig(YoutubeId, "refresh").(float64))
	YoutuberefreshTime *= 120
	Youtuberefresh()
}

func YoutubeOnEnable(bindings *Binding) {

}

func YoutubeOnDisable(bindings *Binding) {

}

func YoutubeOnTick(bindings *Binding) {
	bindings.Clear()
	bindings.Draw(-Youtubecnt, 0, Youtubelogo, 4, 4)
	bindings.DrawString(Youtubecnt-5, Youtubefollowers)
	bindings.Push()

	Youtubecnt++
	if Youtubecnt > 5+DisplayedSize(Youtubefollowers) {
		Youtubecnt = 0
	}

	YoutuberefreshCnt++
	if YoutuberefreshCnt == YoutuberefreshTime {
		Youtuberefresh()
		YoutuberefreshCnt = 0;
	}
}

func Youtuberefresh() {
	doc, err := goquery.NewDocument(GetPluginConfig(YoutubeId, "channel").(string))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".yt-subscription-button-subscriber-count-branded-horizontal").Each(func(i int, s *goquery.Selection) {
		content, err := s.Html()
		if err == nil {
			Youtubefollowers = content
		}
	})
}

var (
	YoutubePluginDescriptor PixxiePlugin = PixxiePlugin{
		YoutubeId, 60, YoutubeOnInit, YoutubeOnEnable, YoutubeOnDisable, YoutubeOnTick,
	}
)
