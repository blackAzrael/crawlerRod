package main

import (
	"crawlerRod/pkg/config"
	"crawlerRod/pkg/logger"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	_config := config.Config
	data, err := json.Marshal(_config)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	//fmt.Println("data:\t", string(data))
	logger.Logger.Debugf(string(data))

}
