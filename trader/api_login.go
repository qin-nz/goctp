package trader

import (
	"errors"
	"time"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

// 用户登录请求
func (p *apispi) ReqUserLogin(reqf libctp.CThostFtdcReqUserLoginField) error {

	time.Sleep(time.Second * 1)

	logrus.WithFields(logrus.Fields{
		"broker_id": reqf.GetBrokerID(),
		"user_id":   reqf.GetUserID(),
	}).Println("交易系统账号登陆中...")

	reqID := p.newRequestId()
	result := p.api.ReqUserLogin(reqf, reqID)

	if result != 0 {
		return errors.New("发送用户登录请求失败！")
	}

	return p.sig.Wait(reqID)

}

func (p *apispi) OnRspUserLogin(pRspUserLogin libctp.CThostFtdcRspUserLoginField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	err := comm.RspToError(pRspInfo)
	if err != nil {
		p.sig.Trigger(nRequestID, err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"trader_day": p.api.GetTradingDay(),
	}).Println("交易账号已登录")
	p.sig.Trigger(nRequestID, nil)
}
