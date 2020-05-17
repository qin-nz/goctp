package md

import (
	"fmt"

	"github.com/qin-nz/goctp/comm"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

// 订阅行情
func (p *apispi) SubscribeMarketData(instrumentIDs ...string) error {

	if len(instrumentIDs) == 0 {
		logrus.Println("没有指定需要订阅的行情数据")
		return nil
	}

	logrus.Println("订阅行情数据中...")

	insPtr := comm.StringSlice2C(instrumentIDs)

	iResult := p.api.SubscribeMarketData(insPtr, len(instrumentIDs))

	if iResult != 0 {
		// TODO:显示错误详情
		return fmt.Errorf("发送订阅行情请求失败！%d", iResult)
	}

	return nil
}

// 订阅行情应答
func (p *apispi) OnRspSubMarketData(pSpecificInstrument libctp.CThostFtdcSpecificInstrumentField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	logrus.WithField("reqid", nRequestID).Info("OnRspSubMarketData")

	if pRspInfo.GetErrorID() != 0 {
		logrus.WithFields(logrus.Fields{
			"errmsg":       pRspInfo.GetErrorMsg(),
			"errid":        pRspInfo.GetErrorID(),
			"InstrumentID": pSpecificInstrument.GetInstrumentID(),
		}).Warn("行情订阅异常")
	}

	logrus.Printf("订阅合约 %#v 行情数据成功！ id %d\n", pSpecificInstrument.GetInstrumentID(), nRequestID)
}

// 深度行情通知
func (p *apispi) OnRtnDepthMarketData(pDepthMarketData libctp.CThostFtdcDepthMarketDataField) {
	logrus.WithFields(logrus.Fields{
		"成交时间": pDepthMarketData.GetUpdateTime(),
		"合约":   pDepthMarketData.GetInstrumentID(),
		"最新价":  pDepthMarketData.GetLastPrice(),
		"买一价":  pDepthMarketData.GetBidPrice1(),
		"卖一价":  pDepthMarketData.GetAskPrice1(),
		"买一量":  pDepthMarketData.GetBidVolume1(),
		"卖一量":  pDepthMarketData.GetAskVolume1(),
	}).Info()
	// TODO: 返回 chan 
}
