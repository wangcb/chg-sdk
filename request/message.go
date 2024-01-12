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
	TempParams     string `json:"temp_params" form:"temp_params"`
	WechatTempId   string `json:"wechat_temp_id" form:"wechat_temp_id"`
	FeishuUrl      string `json:"feishu_url" form:"feishu_url"`
	TargetPlatform string `json:"target_platform" form:"target_platform"`
	Status         int    `json:"status" form:"status"`
	Creator        int    `json:"creator" form:"creator"`
	Updater        int    `json:"updater" form:"updater"`
	Id             int    `json:"id" form:"id"`
}

type SendMessage struct {
	Identify     string `json:"identify" form:"identify" validate:"required" label:"模板标识"`
	PushType     string `json:"push_type" form:"push_type" validate:"required" label:"消息类型"`
	ToUser       string `json:"to_user" form:"to_user"`
	Channel      string `json:"channel" form:"channel" validate:"required" label:"通知渠道 json串"`
	RequestParam string `json:"request_param" form:"request_param" label:"模板参数 json串"`
	Content      string `json:"content" form:"content" label:"消息内容"`
}

type ReadUser struct {
	UserId    int `json:"user_id" validate:"required" label:"用户id"`
	MessageId int `json:"message_id" validate:"required" label:"用户id"`
}
