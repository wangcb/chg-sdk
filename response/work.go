package response

type WorkUserList struct {
	ErrMsg     string `json:"errmsg"`
	NextCursor string `json:"next_cursor"`
	DeptUser   []WorkUserListUser
}

type WorkUserListUser struct {
	UserId     string `json:"userid"`
	Department int64  `json:"department"`
}

type ContactQrcode struct {
	WorkCommon
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

type WorkCommon struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"errmsg"`
}
