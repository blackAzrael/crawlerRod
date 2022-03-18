package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	LogLevel  string `yaml:"log_level"`
	LogPath   string `yaml:"log_path"`
	LogCaller bool   `yaml:"log_caller"`
	User      string `yaml:"user"`
	Pwd       string `yaml:"pwd"`
	Proxy     string `yaml:"proxy"`
}

func (c *conf) GetConf() *conf {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

var Config *conf

func init() {
	var c conf
	Config = c.GetConf()
}
