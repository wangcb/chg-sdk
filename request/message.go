package request

type TemplateList struct {
	Page     int    `json:"page" form:"page"` // 页码
	PageSize int    `json:"size" form:"size"` // 每页大小
	Id       int    `json:"id" form:"id"`
	Appid    string `json:"appid" form:"appid"`
	Channel  string `json:"channel" form:"channel"`
	Status   int    `json:"status" form:"status"`
}

type TemplateParam struct {
	Appid          string `json:"appid" form:"appid" validate:"required" label:"app_id"`
	Identify       string `json:"identify" form:"identify" validate:"required" label:"模板标识"`
	SceneDesc      string `json:"scene_desc" form:"scene_desc" validate:"required" label:"场景说明"`
	Channel        string `json:"channel" form:"channel" validate:"required" label:"通知渠道"`
	Title          string `json:"title" form:"title" validate:"required" label:"标题"`
	Content        string `json:"content" form:"content" validate:"required" label:"content"`
	WechatTempId   string `json:"wechat_temp_id" form:"wechat_temp_id"`
	FeishuUrl      string `json:"feishu_url" form:"feishu_url"`
	TargetPlatform string `json:"target_platform" form:"target_platform"`
	Status         int    `json:"status" form:"status"`
	Creator        int    `json:"creator" form:"creator"`
	Updater        int    `json:"updater" form:"updater"`
	Id             int    `json:"id" form:"id"`
}
