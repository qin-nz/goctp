package md

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
)

type client struct {
	// 行情前置机的地址
	Fronts []string
	md     *apispi
}

func New(mdFront []string) client {
	tempDir, err := ioutil.TempDir("", "md")
	if err != nil {
		logrus.Fatalln("创建临时目录失败")
	}

	api := libctp.CThostFtdcMdApiCreateFtdcMdApi(tempDir)
	as := &apispi{api: api,
		sig: signal.NewManager(),
	}

	api.RegisterSpi(libctp.NewDirectorCThostFtdcMdSpi(as))

	for _, val := range mdFront {
		api.RegisterFront(val)
	}

	return client{md: as, Fronts: mdFront}
}

func (c *client) Init() error {
	return c.md.Init()
}

func (c *client) Login(account comm.Account) error {
	reqf := libctp.NewCThostFtdcReqUserLoginField()
	reqf.SetBrokerID(account.BrokerID)
	reqf.SetUserID(account.InvestorID)
	reqf.SetPassword(account.Password)

	return c.md.ReqUserLogin(reqf)
}

func (c *client) SubscribeMarketData(insts ...string) (chan interface{}, error) {
	err := c.md.SubscribeMarketData(insts...)
	return nil, err
}
