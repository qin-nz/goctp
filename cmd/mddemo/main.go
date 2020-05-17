package main

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/goctp/md"
	"github.com/sirupsen/logrus"
)

func main() {
	config := &Config{}
	configPath := "./conf/app.toml"

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatalf("配置文件错误 %v，路径 %s", err)
	}

	mdc := md.New(config.Md.Front)

	mdc.Init()
	mdc.Login(config.Account)
	mdc.SubscribeMarketData(config.Md.Subscribe...)

	for {
		time.Sleep(time.Minute)
	}
}

type Md struct {
	Front     []string `toml:"front"`
	Subscribe []string `toml:"subscribe"`
}

type Config struct {
	Account comm.Account `toml:"account"`
	Md      Md           `toml:"md"`
}
