package trader

import (
	"errors"

	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

func (p *apispi) Init() error {
	p.api.Init()
	return p.sig.Wait(signal.Init)

	// 进行客户端验证
}

func (p *apispi) OnFrontConnected() {
	logrus.WithFields(logrus.Fields{
		"APIVersion": libctp.CThostFtdcMdApiGetApiVersion(),
	}).Info("交易模块初始化成功")
	p.sig.Trigger(signal.Init, nil)
}

func (p *apispi) OnFrontDisconnected(nReason int) {
	msg := "交易模块初始化失败"
	p.sig.Trigger(signal.Init, errors.New(msg))

	logrus.WithFields(logrus.Fields{
		"reason": nReason,
	}).Warn(msg)
}
