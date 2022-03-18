package main

import (
	"crawlerRod/pkg/config"
	"crawlerRod/pkg/logger"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	path, _ := launcher.LookPath()
	logger.Logger.Info(path)
	proxy := config.Config.Proxy
	logger.Logger.Infof("proxy: %s", proxy)
	u := launcher.New().Bin(path).Proxy(proxy).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://www.google.co.jp/")
	page.MustWaitLoad().MustScreenshot("img/a.png")
	text, _ := page.MustElementX("/html/body/div[1]/div[3]/form/div[1]/div[1]/div[3]/center/input[1]").Text()
	logger.Logger.Debugf(text)

}
