package response

import "time"

type WechatTemplate struct {
	Content              string        `json:"content"`
	TempParams           string        `json:"temp_params"`
	Example              string        `json:"example"`
	KeywordEnumValueList []interface{} `json:"keywordEnumValueList"`
	PriTmplId            string        `json:"priTmplId"`
	Title                string        `json:"title"`
	Type                 int           `json:"type"`
}

type MessageTemplate struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty" mapstructure:"id,squash"`
	AppId          string    `json:"app_id" form:"app_id" gorm:"column:app_id;comment:应用标识;size:32;"`
	Identify       string    `json:"identify" form:"identify" gorm:"column:identify;comment:模板标识;size:32;"`
	SceneDesc      string    `json:"scene_desc" form:"scene_desc" gorm:"column:scene_desc;comment:场景说明;size:60;"`
	Title          string    `json:"title" form:"title" gorm:"column:title;comment:模板标题;size:120;"`
	Content        string    `json:"content" form:"content" gorm:"column:content;comment:模板消息内容;size:255;"`
	TempParams     string    `json:"temp_params" form:"temp_params" gorm:"column:temp_params;comment:模板参数;size:255;"`
	Channel        string    `json:"channel" form:"channel" gorm:"column:channel;comment:通知渠道 wechat:微信 feishu:飞书 getui:个推 web:站内信;size:32;"`
	WechatTempId   string    `json:"wechat_temp_id" form:"wechat_temp_id" gorm:"column:wechat_temp_id;comment:微信模板id;size:120;"`
	FeishuUrl      string    `json:"feishu_url" form:"feishu_url" gorm:"column:feishu_url;comment:飞书通知地址;size:255;"`
	TargetPlatform string    `json:"target_platform" form:"target_platform" gorm:"column:target_platform;comment:目标平台 android,ios;size:80;"`
	GroupType      *int      `json:"group_type" form:"group_type" gorm:"column:group_type;comment:消息分组;size:1;"`
	Status         int       `json:"status" form:"status" gorm:"column:status;comment:1.启用 0.禁用;"`
	Creator        int       `json:"creator" form:"creator" gorm:"column:creator;comment:创建者;size:10;"`
	Updater        int       `json:"updater" form:"updater" gorm:"column:updater;comment:更新者;size:10;"`
	CreatedAt      time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt      time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}

type TemplateList struct {
	TotalCount int64              `json:"totalCount"`
	List       []*MessageTemplate `json:"list"`
}

type Messagelist struct {
	TotalCount int64      `json:"totalCount"`
	List       []*Message `json:"list"`
}

type MessagelistDetail struct {
	Data     *Message       `json:"data"`
	PushData []*MessagePush `json:"pushList"`
}

type Message struct {
	AppId        string    `json:"appId" form:"appId"`
	TemplateId   int       `json:"templateId" form:"templateId"`
	SingleUserId int       `json:"singleUserId" form:"singleUserId"`
	GroupUserId  string    `json:"groupUserId" form:"groupUserId"`
	Title        string    `json:"title" form:"title"`
	Content      string    `json:"content" form:"content"`
	ErrMsg       string    `json:"err_msg" form:"err_msg"`
	PushType     int       `json:"pushType" form:"pushType"`
	GroupType    int       `json:"group_type" form:"group_type"`
	CreatedAt    time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
}

type MessagePush struct {
	MessageId    int    `json:"messageId" form:"messageId"`
	Channel      string `json:"channel" form:"channel"`
	Status       int    `json:"status" form:"status"`
	RequestParam string `json:"requestParam" form:"requestParam"`
	Response     string `json:"response" form:"response"`
	Reason       string `json:"reason" form:"reason"`
}
