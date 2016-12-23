package duobb_proto

const (
	AccountService = "DuobbAccountService"
	SPPlanService  = "SelectProductService"
)

const (
	DUOBB_RSP_SUCCESS = iota
	DUOBB_RSP_ERROR
	DUOBB_MSG_GET_ACCOUNT_ERROR
	DUOBB_MSG_DECODE_ERROR
	DUOBB_MSG_LOGIN_ERROR
	DUOBB_MSG_LOGOUT_ERROR
	DUOBB_MSG_PROCESS_ERROR
	DUOBB_DB_ERROR
	DUOBB_GET_PLAN_FROM_PW_ERROR
)

const (
	MSG_DUOBB_DB_ERROR           = "数据库执行失败"
	MSG_DUOBB_LOGIN_ERROR        = "登录失败"
	MSG_DUOBB_LOGOUT_ERROR       = "登出失败"
	MSG_DUOBB_PLAN_FROM_PW_ERROR = "密码获取选品计划失败"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
