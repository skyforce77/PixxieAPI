package main

import (
	"log"
	"github.com/PuerkitoBio/goquery"
)

var (
	Twitterc1 = Pixel{0x55, 0x84, 0xba}
	Twitterc2 = Pixel{0xc9, 0xca, 0x61}
	Twitterc3 = Pixel{0x00, 0x00, 0x00}

	Twitterlogo = []Pixel{
		Twitterc3, Twitterc3, Twitterc3, Twitterc1,
		Twitterc1, Twitterc3, Twitterc1, Twitterc2,
		Twitterc1, Twitterc1, Twitterc1, Twitterc1,
		Twitterc3, Twitterc1, Twitterc1, Twitterc3,
	}

	TwitterId = "twitter"

	Twittercnt = 0
	Twitterfollowers = "0"

	TwitterrefreshTime = 0
	TwitterrefreshCnt = 0
)

func TwitterOnInit() {
	TwitterrefreshTime = int(GetPluginConfig(TwitterId, "refresh").(float64))
	TwitterrefreshTime *= 120
	Twitterrefresh()
}

func TwitterOnEnable(bindings *Binding) {

}

func TwitterOnDisable(bindings *Binding) {

}

func TwitterOnTick(bindings *Binding) {
	bindings.Clear()
	bindings.Draw(-Twittercnt, 0, Twitterlogo, 4, 4)
	bindings.DrawString(Twittercnt-5, Twitterfollowers)
	bindings.Push()

	Twittercnt++
	if Twittercnt > 5+DisplayedSize(Twitterfollowers) {
		Twittercnt = 0
	}

	TwitterrefreshCnt++
	if TwitterrefreshCnt == TwitterrefreshTime {
		Twitterrefresh()
		TwitterrefreshCnt = 0;
	}
}

func Twitterrefresh() {
	doc, err := goquery.NewDocument(GetPluginConfig(TwitterId, "user").(string))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".ProfileNav-item--followers > .ProfileNav-stat > .ProfileNav-value").Each(func(i int, s *goquery.Selection) {
		content, err := s.Html()
		if err == nil {
			Twitterfollowers = content
		}
	})
}

var (
	TwitterPluginDescriptor PixxiePlugin = PixxiePlugin{
		TwitterId, 60, TwitterOnInit, TwitterOnEnable, TwitterOnDisable, TwitterOnTick,
	}
)
