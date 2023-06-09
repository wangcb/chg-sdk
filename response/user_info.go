package response

type UserInfo struct {
	Id         int64  `json:"id"`
	NickName   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	RegSource  int    `json:"reg_source"`
	IsRealName int    `json:"is_real_name"`
}
