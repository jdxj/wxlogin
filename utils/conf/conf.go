package conf

import (
	"gopkg.in/ini.v1"
)

const (
	configPath = "config.ini"
)

var (
	Config *ini.File
)

func init() {
	conf, err := ini.Load(configPath)
	if err != nil {
		panic(err)
	}
	Config = conf
}
