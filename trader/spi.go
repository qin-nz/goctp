package md

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/qin-nz/libctp"
)

type CThostFtdcMdSpiLogin interface {

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

	///客户端认证响应
	OnRspAuthenticate(pRspAuthenticateField libctp.CThostFtdcRspAuthenticateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///登录请求响应
	OnRspUserLogin(pRspUserLogin libctp.CThostFtdcRspUserLoginField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///登出请求响应
	OnRspUserLogout(pUserLogout libctp.CThostFtdcUserLogoutField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///用户口令更新请求响应
	OnRspUserPasswordUpdate(pUserPasswordUpdate libctp.CThostFtdcUserPasswordUpdateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///资金账户口令更新请求响应
	OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate libctp.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///查询用户当前支持的认证模式的回复
	OnRspUserAuthMethod(pRspUserAuthMethod libctp.CThostFtdcRspUserAuthMethodField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///获取图形验证码请求的回复
	OnRspGenUserCaptcha(pRspGenUserCaptcha libctp.CThostFtdcRspGenUserCaptchaField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///获取短信验证码请求的回复
	OnRspGenUserText(pRspGenUserText libctp.CThostFtdcRspGenUserTextField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)
}

type CThostFtdcMdSpi interface {
	CThostFtdcMdSpiLogin

	///心跳超时警告。当长时间未收到报文时，该方法被调用。
	///@param nTimeLapse 距离上次接收报文的时间
	OnHeartBeatWarning(nTimeLapse int)

	///报单录入请求响应
	OnRspOrderInsert(pInputOrder libctp.CThostFtdcInputOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///预埋单录入请求响应
	OnRspParkedOrderInsert(pParkedOrder libctp.CThostFtdcParkedOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///预埋撤单录入请求响应
	OnRspParkedOrderAction(pParkedOrderAction libctp.CThostFtdcParkedOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///报单操作请求响应
	OnRspOrderAction(pInputOrderAction libctp.CThostFtdcInputOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///查询最大报单数量响应
	OnRspQueryMaxOrderVolume(pQueryMaxOrderVolume libctp.CThostFtdcQueryMaxOrderVolumeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///投资者结算结果确认响应
	OnRspSettlementInfoConfirm(pSettlementInfoConfirm libctp.CThostFtdcSettlementInfoConfirmField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///删除预埋单响应
	OnRspRemoveParkedOrder(pRemoveParkedOrder libctp.CThostFtdcRemoveParkedOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///删除预埋撤单响应
	OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction libctp.CThostFtdcRemoveParkedOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///执行宣告录入请求响应
	OnRspExecOrderInsert(pInputExecOrder libctp.CThostFtdcInputExecOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///执行宣告操作请求响应
	OnRspExecOrderAction(pInputExecOrderAction libctp.CThostFtdcInputExecOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///询价录入请求响应
	OnRspForQuoteInsert(pInputForQuote libctp.CThostFtdcInputForQuoteField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///报价录入请求响应
	OnRspQuoteInsert(pInputQuote libctp.CThostFtdcInputQuoteField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///报价操作请求响应
	OnRspQuoteAction(pInputQuoteAction libctp.CThostFtdcInputQuoteActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///批量报单操作请求响应
	OnRspBatchOrderAction(pInputBatchOrderAction libctp.CThostFtdcInputBatchOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///期权自对冲录入请求响应
	OnRspOptionSelfCloseInsert(pInputOptionSelfClose libctp.CThostFtdcInputOptionSelfCloseField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///期权自对冲操作请求响应
	OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction libctp.CThostFtdcInputOptionSelfCloseActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///申请组合录入请求响应
	OnRspCombActionInsert(pInputCombAction libctp.CThostFtdcInputCombActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询报单响应
	OnRspQryOrder(pOrder libctp.CThostFtdcOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询成交响应
	OnRspQryTrade(pTrade libctp.CThostFtdcTradeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资者持仓响应
	OnRspQryInvestorPosition(pInvestorPosition libctp.CThostFtdcInvestorPositionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询资金账户响应
	OnRspQryTradingAccount(pTradingAccount libctp.CThostFtdcTradingAccountField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资者响应
	OnRspQryInvestor(pInvestor libctp.CThostFtdcInvestorField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询交易编码响应
	OnRspQryTradingCode(pTradingCode libctp.CThostFtdcTradingCodeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询合约保证金率响应
	OnRspQryInstrumentMarginRate(pInstrumentMarginRate libctp.CThostFtdcInstrumentMarginRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询合约手续费率响应
	OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate libctp.CThostFtdcInstrumentCommissionRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询交易所响应
	OnRspQryExchange(pExchange libctp.CThostFtdcExchangeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询产品响应
	OnRspQryProduct(pProduct libctp.CThostFtdcProductField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询合约响应
	OnRspQryInstrument(pInstrument libctp.CThostFtdcInstrumentField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询行情响应
	OnRspQryDepthMarketData(pDepthMarketData libctp.CThostFtdcDepthMarketDataField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资者结算结果响应
	OnRspQrySettlementInfo(pSettlementInfo libctp.CThostFtdcSettlementInfoField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询转帐银行响应
	OnRspQryTransferBank(pTransferBank libctp.CThostFtdcTransferBankField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资者持仓明细响应
	OnRspQryInvestorPositionDetail(pInvestorPositionDetail libctp.CThostFtdcInvestorPositionDetailField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询客户通知响应
	OnRspQryNotice(pNotice libctp.CThostFtdcNoticeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询结算信息确认响应
	OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm libctp.CThostFtdcSettlementInfoConfirmField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资者持仓明细响应
	OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail libctp.CThostFtdcInvestorPositionCombineDetailField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///查询保证金监管系统经纪公司资金账户密钥响应
	OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey libctp.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询仓单折抵信息响应
	OnRspQryEWarrantOffset(pEWarrantOffset libctp.CThostFtdcEWarrantOffsetField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资者品种/跨品种保证金响应
	OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin libctp.CThostFtdcInvestorProductGroupMarginField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询交易所保证金率响应
	OnRspQryExchangeMarginRate(pExchangeMarginRate libctp.CThostFtdcExchangeMarginRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询交易所调整保证金率响应
	OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust libctp.CThostFtdcExchangeMarginRateAdjustField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询汇率响应
	OnRspQryExchangeRate(pExchangeRate libctp.CThostFtdcExchangeRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询二级代理操作员银期权限响应
	OnRspQrySecAgentACIDMap(pSecAgentACIDMap libctp.CThostFtdcSecAgentACIDMapField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询产品报价汇率
	OnRspQryProductExchRate(pProductExchRate libctp.CThostFtdcProductExchRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询产品组
	OnRspQryProductGroup(pProductGroup libctp.CThostFtdcProductGroupField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询做市商合约手续费率响应
	OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate libctp.CThostFtdcMMInstrumentCommissionRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询做市商期权合约手续费响应
	OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate libctp.CThostFtdcMMOptionInstrCommRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询报单手续费响应
	OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate libctp.CThostFtdcInstrumentOrderCommRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询资金账户响应
	OnRspQrySecAgentTradingAccount(pTradingAccount libctp.CThostFtdcTradingAccountField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询二级代理商资金校验模式响应
	OnRspQrySecAgentCheckMode(pSecAgentCheckMode libctp.CThostFtdcSecAgentCheckModeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询二级代理商信息响应
	OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo libctp.CThostFtdcSecAgentTradeInfoField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询期权交易成本响应
	OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost libctp.CThostFtdcOptionInstrTradeCostField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询期权合约手续费响应
	OnRspQryOptionInstrCommRate(pOptionInstrCommRate libctp.CThostFtdcOptionInstrCommRateField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询执行宣告响应
	OnRspQryExecOrder(pExecOrder libctp.CThostFtdcExecOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询询价响应
	OnRspQryForQuote(pForQuote libctp.CThostFtdcForQuoteField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询报价响应
	OnRspQryQuote(pQuote libctp.CThostFtdcQuoteField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询期权自对冲响应
	OnRspQryOptionSelfClose(pOptionSelfClose libctp.CThostFtdcOptionSelfCloseField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询投资单元响应
	OnRspQryInvestUnit(pInvestUnit libctp.CThostFtdcInvestUnitField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询组合合约安全系数响应
	OnRspQryCombInstrumentGuard(pCombInstrumentGuard libctp.CThostFtdcCombInstrumentGuardField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询申请组合响应
	OnRspQryCombAction(pCombAction libctp.CThostFtdcCombActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询转帐流水响应
	OnRspQryTransferSerial(pTransferSerial libctp.CThostFtdcTransferSerialField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询银期签约关系响应
	OnRspQryAccountregister(pAccountregister libctp.CThostFtdcAccountregisterField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///错误应答
	OnRspError(pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///报单通知
	OnRtnOrder(pOrder libctp.CThostFtdcOrderField)

	///成交通知
	OnRtnTrade(pTrade libctp.CThostFtdcTradeField)

	///报单录入错误回报
	OnErrRtnOrderInsert(pInputOrder libctp.CThostFtdcInputOrderField, pRspInfo libctp.CThostFtdcRspInfoField)

	///报单操作错误回报
	OnErrRtnOrderAction(pOrderAction libctp.CThostFtdcOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField)

	///合约交易状态通知
	OnRtnInstrumentStatus(pInstrumentStatus libctp.CThostFtdcInstrumentStatusField)

	///交易所公告通知
	OnRtnBulletin(pBulletin libctp.CThostFtdcBulletinField)

	///交易通知
	OnRtnTradingNotice(pTradingNoticeInfo libctp.CThostFtdcTradingNoticeInfoField)

	///提示条件单校验错误
	OnRtnErrorConditionalOrder(pErrorConditionalOrder libctp.CThostFtdcErrorConditionalOrderField)

	///执行宣告通知
	OnRtnExecOrder(pExecOrder libctp.CThostFtdcExecOrderField)

	///执行宣告录入错误回报
	OnErrRtnExecOrderInsert(pInputExecOrder libctp.CThostFtdcInputExecOrderField, pRspInfo libctp.CThostFtdcRspInfoField)

	///执行宣告操作错误回报
	OnErrRtnExecOrderAction(pExecOrderAction libctp.CThostFtdcExecOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField)

	///询价录入错误回报
	OnErrRtnForQuoteInsert(pInputForQuote libctp.CThostFtdcInputForQuoteField, pRspInfo libctp.CThostFtdcRspInfoField)

	///报价通知
	OnRtnQuote(pQuote libctp.CThostFtdcQuoteField)

	///报价录入错误回报
	OnErrRtnQuoteInsert(pInputQuote libctp.CThostFtdcInputQuoteField, pRspInfo libctp.CThostFtdcRspInfoField)

	///报价操作错误回报
	OnErrRtnQuoteAction(pQuoteAction libctp.CThostFtdcQuoteActionField, pRspInfo libctp.CThostFtdcRspInfoField)

	///询价通知
	OnRtnForQuoteRsp(pForQuoteRsp libctp.CThostFtdcForQuoteRspField)

	///保证金监控中心用户令牌
	OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken libctp.CThostFtdcCFMMCTradingAccountTokenField)

	///批量报单操作错误回报
	OnErrRtnBatchOrderAction(pBatchOrderAction libctp.CThostFtdcBatchOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField)

	///期权自对冲通知
	OnRtnOptionSelfClose(pOptionSelfClose libctp.CThostFtdcOptionSelfCloseField)

	///期权自对冲录入错误回报
	OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose libctp.CThostFtdcInputOptionSelfCloseField, pRspInfo libctp.CThostFtdcRspInfoField)

	///期权自对冲操作错误回报
	OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction libctp.CThostFtdcOptionSelfCloseActionField, pRspInfo libctp.CThostFtdcRspInfoField)

	///申请组合通知
	OnRtnCombAction(pCombAction libctp.CThostFtdcCombActionField)

	///申请组合录入错误回报
	OnErrRtnCombActionInsert(pInputCombAction libctp.CThostFtdcInputCombActionField, pRspInfo libctp.CThostFtdcRspInfoField)

	///请求查询签约银行响应
	OnRspQryContractBank(pContractBank libctp.CThostFtdcContractBankField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询预埋单响应
	OnRspQryParkedOrder(pParkedOrder libctp.CThostFtdcParkedOrderField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询预埋撤单响应
	OnRspQryParkedOrderAction(pParkedOrderAction libctp.CThostFtdcParkedOrderActionField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询交易通知响应
	OnRspQryTradingNotice(pTradingNotice libctp.CThostFtdcTradingNoticeField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询经纪公司交易参数响应
	OnRspQryBrokerTradingParams(pBrokerTradingParams libctp.CThostFtdcBrokerTradingParamsField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询经纪公司交易算法响应
	OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos libctp.CThostFtdcBrokerTradingAlgosField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///请求查询监控中心用户令牌
	OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken libctp.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///银行发起银行资金转期货通知
	OnRtnFromBankToFutureByBank(pRspTransfer libctp.CThostFtdcRspTransferField)

	///银行发起期货资金转银行通知
	OnRtnFromFutureToBankByBank(pRspTransfer libctp.CThostFtdcRspTransferField)

	///银行发起冲正银行转期货通知
	OnRtnRepealFromBankToFutureByBank(pRspRepeal libctp.CThostFtdcRspRepealField)

	///银行发起冲正期货转银行通知
	OnRtnRepealFromFutureToBankByBank(pRspRepeal libctp.CThostFtdcRspRepealField)

	///期货发起银行资金转期货通知
	OnRtnFromBankToFutureByFuture(pRspTransfer libctp.CThostFtdcRspTransferField)

	///期货发起期货资金转银行通知
	OnRtnFromFutureToBankByFuture(pRspTransfer libctp.CThostFtdcRspTransferField)

	///系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal libctp.CThostFtdcRspRepealField)

	///系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal libctp.CThostFtdcRspRepealField)

	///期货发起查询银行余额通知
	OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount libctp.CThostFtdcNotifyQueryAccountField)

	///期货发起银行资金转期货错误回报
	OnErrRtnBankToFutureByFuture(pReqTransfer libctp.CThostFtdcReqTransferField, pRspInfo libctp.CThostFtdcRspInfoField)

	///期货发起期货资金转银行错误回报
	OnErrRtnFutureToBankByFuture(pReqTransfer libctp.CThostFtdcReqTransferField, pRspInfo libctp.CThostFtdcRspInfoField)

	///系统运行时期货端手工发起冲正银行转期货错误回报
	OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal libctp.CThostFtdcReqRepealField, pRspInfo libctp.CThostFtdcRspInfoField)

	///系统运行时期货端手工发起冲正期货转银行错误回报
	OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal libctp.CThostFtdcReqRepealField, pRspInfo libctp.CThostFtdcRspInfoField)

	///期货发起查询银行余额错误回报
	OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount libctp.CThostFtdcReqQueryAccountField, pRspInfo libctp.CThostFtdcRspInfoField)

	///期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFuture(pRspRepeal libctp.CThostFtdcRspRepealField)

	///期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFuture(pRspRepeal libctp.CThostFtdcRspRepealField)

	///期货发起银行资金转期货应答
	OnRspFromBankToFutureByFuture(pReqTransfer libctp.CThostFtdcReqTransferField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///期货发起期货资金转银行应答
	OnRspFromFutureToBankByFuture(pReqTransfer libctp.CThostFtdcReqTransferField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///期货发起查询银行余额应答
	OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount libctp.CThostFtdcReqQueryAccountField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLas bool)

	///银行发起银期开户通知
	OnRtnOpenAccountByBank(pOpenAccount libctp.CThostFtdcOpenAccountField)

	///银行发起银期销户通知
	OnRtnCancelAccountByBank(pCancelAccount libctp.CThostFtdcCancelAccountField)

	///银行发起变更银行账号通知
	OnRtnChangeAccountByBank(pChangeAccount libctp.CThostFtdcChangeAccountField)
}
