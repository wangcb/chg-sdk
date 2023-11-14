package response

type UserWechatInfo struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	AppId   string `json:"app_id"`
	OpenId  string `json:"open_id"`
	UnionId string `json:"union_id"`
	AppType string `json:"app_type"`
}
