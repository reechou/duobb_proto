package duobb_proto

const (
	DUOBB_RSP_SUCCESS = iota
	DUOBB_RSP_ERROR
	DUOBB_MSG_GET_ACCOUNT_ERROR
	DUOBB_MSG_DECODE_ERROR
	DUOBB_MSG_LOGIN_ERROR
	DUOBB_DB_ERROR
)

const (
	MSG_DUOBB_DB_ERROR    = "数据库执行失败"
	MSG_DUOBB_LOGIN_ERROR = "登录失败"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
