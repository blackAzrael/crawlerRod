package main

import (
	"crawlerRod/pkg/logger"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestProxyHttp(t *testing.T) {
	uri, err := url.Parse("socks5://127.0.0.1:1080")

	if err != nil {
		log.Fatal("parse url error: ", err)
	}
	client := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}
	resp, err := client.Get("https://www.google.co.jp/")
	if err != nil {
		fmt.Println(err.Error())
	}
	pageBytes, err := ioutil.ReadAll(resp.Body)
	//pageStr := string(pageBytes)
	logger.Logger.Debugf(string(pageBytes))
}
