package main

import (
	"github.com/BurntSushi/toml"
	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/goctp/trader"
	"github.com/sirupsen/logrus"
)

func main() {
	config := &Config{}
	configPath := "./conf/app.toml"

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatalf("配置文件错误 %v，路径 %s", err)
	}

	tr := trader.New(config.Trader.Front)

	tr.Init()
	tr.Auth(config.ClientAuth, config.Account)
	tr.Login(config.Account)
	tr.QrySettlementInfo()
	tr.SettlementInfoConfirm()

}

type Trader struct {
	Front     []string `toml:"front"`
	Subscribe []string `toml:"subscribe"`
}

type Config struct {
	ClientAuth comm.ClientAuth `toml:"client_auth"`
	Account    comm.Account    `toml:"account"`
	Trader     Trader          `toml:"trader"`
}
