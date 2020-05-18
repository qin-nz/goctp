package trader

import (
	"sync"

	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
)

type apispi struct {
	// 请求编号
	requestId int
	idMutex   sync.Mutex

	api libctp.CThostFtdcTraderApi

	// signal
	sig signal.SignalManager
}

func (p *apispi) newRequestId() int {
	p.idMutex.Lock()
	p.requestId++
	p.idMutex.Unlock()
	return p.requestId
}
