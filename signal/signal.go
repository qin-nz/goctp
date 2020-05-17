// signal 用于将异步消息转化为同步消息使用
package signal

import "github.com/sirupsen/logrus"

var Init int = -1
var Login int = -2

func NewManager() SignalManager {
	return &signalManager{
		c: make(map[int]chan error),
	}
}

type SignalManager interface {
	Wait(id int) error
	Trigger(id int, err error)
}

type signalManager struct {
	c map[int]chan error
}

func (m *signalManager) Wait(id int) error {
	m.c[id] = make(chan error)
	return <-m.c[id]
}

func (m *signalManager) Trigger(id int, err error) {
	if id == 0 {
		logrus.Debug("RequestID=0")
		return
	}
	m.c[id] <- err
	delete(m.c, id)
}
