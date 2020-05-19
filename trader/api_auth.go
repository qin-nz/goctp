package trader

import (
	"errors"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

// 客户端认证
func (p *apispi) ReqAuthenticate(reqf libctp.CThostFtdcReqAuthenticateField) error {

	logrus.WithFields(logrus.Fields{
		"broker_id": reqf.GetBrokerID(),
		"user_id":   reqf.GetUserID(),
		"app_id":    reqf.GetAppID(),
	}).Println("交易系统客户端认证")

	reqID := p.newRequestId()

	result := p.api.ReqAuthenticate(reqf, reqID)

	if result != 0 {
		return errors.New("客户端认证失败")
	}
	return p.sig.Wait(reqID)
}

// 客户端认证响应
func (p *apispi) OnRspAuthenticate(pRspAuthenticateField libctp.CThostFtdcRspAuthenticateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	err := comm.RspToError(pRspInfo)
	if err != nil {
		p.sig.Trigger(nRequestID, err)
		return
	}

	p.sig.Trigger(nRequestID, err)
	logrus.Println("客户端认证成功！")
}
