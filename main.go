package main

import (
	"crawlerRod/pkg/logger"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

)

func main() {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://www.baidu.com")
	page.MustWaitLoad().MustScreenshot("img/a.png")
	text, _ := page.MustElementX("//a").Text()
	logger.Logger.Debugf(text)

}
