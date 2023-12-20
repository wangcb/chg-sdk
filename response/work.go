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
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ContactList struct {
	WorkCommon
	ContactWay []ContactListItem `json:"contact_way"`
}

type ContactListItem struct {
	ConfigId string `json:"config_id"`
}

type ContactDetail struct {
	WorkCommon
	ContactWay map[string]interface{} `json:"contact_way"`
}

type WorkDepartmentUsers struct {
	WorkCommon
	UserList []WorkDepartmentUserList `json:"userlist"`
}

type WorkDepartmentUserList struct {
	UserId     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	OpenUserid string `json:"open_userid"`
}

type WorkDepartmentResp struct {
	WorkCommon
	DepartmentList []WorkDepartment `json:"department"`
}

type WorkDepartment struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentId         int      `json:"parentid"`
	Order            int      `json:"order"`
}
