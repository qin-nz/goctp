package md

import (
	"errors"

	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

func (p *apispi) Init() error {
	p.api.Init()
	return p.s.Wait(signal.Init)
}

func (p *apispi) OnFrontConnected() {
	logrus.WithField("APIVersion", libctp.CThostFtdcMdApiGetApiVersion()).Info("行情模块初始化成功")
	p.s.Trigger(signal.Init, nil)
}

func (p *apispi) OnFrontDisconnected(nReason int) {
	msg := "行情模块初始化失败"
	p.s.Trigger(signal.Init, errors.New(msg))
	logrus.WithField("reason", nReason).Warn(msg)
}
