package response

type UserInfo struct {
	Id         int64  `json:"id"`
	NickName   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	RegSource  int    `json:"reg_source"`
	IsRealName int    `json:"is_real_name"`
	Activity   string `json:"activity"`
}

type WorkExternalUser struct {
	UserInfo
	WorkUserId string `json:"work_user_id"`
}

type WorkExternalUserListByCore struct {
	List  []WorkExternalUser `json:"list"`
	Total int64              `json:"total"`
}
