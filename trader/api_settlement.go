package trader

import (
	"errors"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

func (p *apispi) ReqQrySettlementInfo(set libctp.CThostFtdcQrySettlementInfoField) error {
	reqID := p.newRequestId()

	result := p.api.ReqQrySettlementInfo(set, reqID)
	comm.LogReq(reqID, "查询结算单", result, set)

	if result != 0 {
		return errors.New("结算单查询失败！")
	}

	return p.sig.Wait(reqID)
}

///请求查询投资者结算结果响应
func (p *apispi) OnRspQrySettlementInfo(pSettlementInfo libctp.CThostFtdcSettlementInfoField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	var err error
	defer p.sig.Trigger(nRequestID, err)

	err = comm.RspToError(pRspInfo)
	if err != nil {
		return
	}

	if uintptr(pSettlementInfo.(libctp.SwigcptrCThostFtdcSettlementInfoField)) == 0 {
		logrus.WithFields(logrus.Fields{
			"reqID": nRequestID,
		}).Warn("OnRspQrySettlementInfo pSettlementInfo=null")
		return
	}

	logrus.WithFields(logrus.Fields{
		"content": comm.ShouldDecodeGBK(pSettlementInfo.GetContent()),
	}).Println("查询投资者结算单")
}

// 投资者结算单确认
func (p *apispi) ReqSettlementInfoConfirm(confirm libctp.CThostFtdcSettlementInfoConfirmField) error {

	reqID := p.newRequestId()

	result := p.api.ReqSettlementInfoConfirm(confirm, reqID)
	comm.LogReq(reqID, "确认结算单", result, confirm)

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

	logrus.WithFields(logrus.Fields{}).Println("结算单已确认")
	p.sig.Trigger(nRequestID, nil)
}
