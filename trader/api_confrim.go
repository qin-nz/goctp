package trader

import (
	"errors"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

// 投资者结算单确认
func (p *apispi) ReqSettlementInfoConfirm(confirm libctp.CThostFtdcSettlementInfoConfirmField) error {

	reqID := p.newRequestId()

	result := p.api.ReqSettlementInfoConfirm(confirm, reqID)

	if result != 0 {
		return errors.New("确认投资者结算单失败！")
	}
	return p.sig.Wait(reqID)

}

// 发送投资者结算单确认响应
func (p *apispi) OnRspSettlementInfoConfirm(pSettlementInfoConfirm libctp.CThostFtdcSettlementInfoConfirmField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	err := comm.RspToError(pRspInfo)
	if err != nil {
		p.sig.Trigger(nRequestID, err)
		return
	}

	logrus.WithFields(logrus.Fields{}).Println("确认投资者结算单")
	p.sig.Trigger(nRequestID, nil)
}
