package comm

// Account 为交易CTP交易账户，可用于登录行情、交易前置
type Account struct {
	BrokerID   string `toml:"broker_id"`
	InvestorID string `toml:"investor_id"`
	Password   string `toml:"password"`
}
