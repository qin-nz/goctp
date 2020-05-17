package md

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"sync"

	"github.com/qin-nz/libctp"

	"github.com/qin-nz/goctp/signal"
)

type apispi struct {
	// 行情请求编号
	requestId int
	idMutex   sync.Mutex

	api libctp.CThostFtdcMdApi

	// signal
	s signal.SignalManager
}

func (p *apispi) newRequestId() int {
	p.idMutex.Lock()
	p.requestId++
	p.idMutex.Unlock()
	return p.requestId
}
