package comm

// Account 为交易CTP交易账户，可用于登录行情、交易前置
type Account struct {
	BrokerID string `toml:"broker_id"`
	UserID   string `toml:"user_id"`
	Password string `toml:"password"`
}

type ClientAuth struct {
	AppID    string `toml:"app_id"`
	AuthCode string `toml:"auth_code"`
}
