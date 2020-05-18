package comm

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

func RspToError(pRspInfo libctp.CThostFtdcRspInfoField) error {
	// 获取调用者信息
	fpcs := make([]uintptr, 1)
	runtime.Callers(2, fpcs)
	caller := runtime.FuncForPC(fpcs[0])
	file, line := caller.FileLine(fpcs[0])

	if pRspInfo.GetErrorID() != 0 {
		logrus.WithFields(logrus.Fields{
			"func":     caller.Name(),
			"location": fmt.Sprintf("%s:%d", file, line),
			"errmsg":   ShouldDecodeGBK(pRspInfo.GetErrorMsg()),
			"errid":    pRspInfo.GetErrorID(),
		}).Warn("CTP返回异常")

		return errors.New(caller.Name() + " CTP返回异常 " + ShouldDecodeGBK(pRspInfo.GetErrorMsg()))
	}
	return nil

}
