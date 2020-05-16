package md

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/qin-nz/libctp"
)

type CThostFtdcMdSpiBasic interface {
	///当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontConnected()

	///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	///@param nReason 错误原因
	///        0x1001 网络读失败
	///        0x1002 网络写失败
	///        0x2001 接收心跳超时
	///        0x2002 发送心跳失败
	///        0x2003 收到错误报文
	OnFrontDisconnected(nReason int)

	///心跳超时警告。当长时间未收到报文时，该方法被调用。
	///@param nTimeLapse 距离上次接收报文的时间
	OnHeartBeatWarning(nTimeLapse int)
}

type CThostFtdcMdSpiLogin interface {
	///登录请求响应
	OnRspUserLogin(pRspUserLogin libctp.CThostFtdcRspUserLoginField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	///登出请求响应
	OnRspUserLogout(pUserLogout libctp.CThostFtdcUserLogoutField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
}

type CThostFtdcMdSpiMarket interface {
	///订阅行情应答
	OnRspSubMarketData(pSpecificInstrument libctp.CThostFtdcSpecificInstrumentField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	///取消订阅行情应答
	OnRspUnSubMarketData(pSpecificInstrument libctp.CThostFtdcSpecificInstrumentField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	///深度行情通知
	OnRtnDepthMarketData(pDepthMarketData libctp.CThostFtdcDepthMarketDataField)
}

type CThostFtdcMdSpiQuote interface {
	///订阅询价应答
	OnRspSubForQuoteRsp(pSpecificInstrument libctp.CThostFtdcSpecificInstrumentField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///取消订阅询价应答
	OnRspUnSubForQuoteRsp(pSpecificInstrument libctp.CThostFtdcSpecificInstrumentField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///询价通知
	OnRtnForQuoteRsp(pForQuoteRsp libctp.CThostFtdcForQuoteRspField)
}

type CThostFtdcMdSpiError interface {
	///错误应答
	OnRspError(pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
}

type CThostFtdcMdSpi interface {
	CThostFtdcMdSpiBasic
	CThostFtdcMdSpiLogin
	CThostFtdcMdSpiMarket
	CThostFtdcMdSpiQuote
	CThostFtdcMdSpiError
}
