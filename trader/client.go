package trader

import (
	"io/ioutil"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

type client struct {
	fronts []string
	tr     *apispi
}

func New(traderFronts []string) client {
	tempDir, err := ioutil.TempDir("", "md")
	if err != nil {
		logrus.Fatalln("创建临时目录失败")
	}

	api := libctp.CThostFtdcTraderApiCreateFtdcTraderApi(tempDir)
	as := &apispi{
		api: api,
		sig: signal.NewManager(),
	}

	api.RegisterSpi(libctp.NewDirectorCThostFtdcTraderSpi(as))

	for _, front := range traderFronts {
		api.RegisterFront(front)
	}

	api.SubscribePublicTopic(libctp.THOST_TERT_RESTART)
	api.SubscribePrivateTopic(libctp.THOST_TERT_RESTART)

	return client{tr: as, fronts: traderFronts}

}

func (c *client) Init() error {
	return c.tr.Init()
}

func (c *client) Login(account comm.Account) error {
	reqf := libctp.NewCThostFtdcReqUserLoginField()
	reqf.SetBrokerID(account.BrokerID)
	reqf.SetUserID(account.InvestorID)
	reqf.SetPassword(account.Password)

	return c.tr.ReqUserLogin(reqf)
}
