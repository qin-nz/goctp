package trader

import (
	"io/ioutil"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

type client struct {
	fronts  []string
	tr      *apispi
	account comm.Account
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

func (c *client) Auth(auth comm.ClientAuth, account comm.Account) error {
	req := libctp.NewCThostFtdcReqAuthenticateField()
	req.SetBrokerID(account.BrokerID)
	req.SetUserID(account.UserID)
	req.SetAppID(auth.AppID)
	req.SetAuthCode(auth.AuthCode)

	return c.tr.ReqAuthenticate(req)
}

func (c *client) QrySettlementInfo() error {
	f := libctp.NewCThostFtdcQrySettlementInfoField()
	f.SetAccountID(c.account.UserID)
	f.SetBrokerID(c.account.BrokerID)
	return c.tr.ReqQrySettlementInfo(f)
}

func (c *client) SettlementInfoConfirm() error {
	confirm := libctp.NewCThostFtdcSettlementInfoConfirmField()
	confirm.SetBrokerID(c.account.BrokerID)
	confirm.SetInvestorID(c.account.UserID) //TOOD.InvestorID)
	return c.tr.ReqSettlementInfoConfirm(confirm)
}

func (c *client) Login(account comm.Account) error {
	reqf := libctp.NewCThostFtdcReqUserLoginField()
	reqf.SetBrokerID(account.BrokerID)
	reqf.SetUserID(account.UserID)
	reqf.SetPassword(account.Password)

	err := c.tr.ReqUserLogin(reqf)
	if err == nil {
		c.account = account
	}
	return err
}
