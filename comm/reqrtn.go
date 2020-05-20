package comm

//go:generate stringer -type=RequestReturnValue

type RequestReturnValue int

const (
	OK                   RequestReturnValue = 0  // 0，代表成功。
	NetworkError                            = -1 //-1，表示网络连接失败；
	UnhandlerExceedLimit                    = -2 //-2，表示未处理请求超过许可数；
	QPSExceedLimit                          = -3 //-3，表示每秒发送请求数超过许可数。
)
