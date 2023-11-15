package response

type ResponseSendSucscribeMessage struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	Msg        string `json:"msg,omitempty"` // 小程序直播的部分接口会把错误提示抛在msg字段
	ErrCode    int    `json:"errcode,omitempty"`
	ErrMsg     string `json:"errmsg,omitempty"`
	ResultCode string `json:"resultcode,omitempty"`
	ResultMsg  string `json:"resultmsg,omitempty"`
}
