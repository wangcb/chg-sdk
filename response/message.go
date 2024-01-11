package response

import "time"

type WechatTemplate struct {
	Content              string        `json:"content"`
	Example              string        `json:"example"`
	KeywordEnumValueList []interface{} `json:"keywordEnumValueList"`
	PriTmplId            string        `json:"priTmplId"`
	Title                string        `json:"title"`
	Type                 int           `json:"type"`
}

type MessageTemplate struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty" mapstructure:"id,squash"`
	AppId          string    `json:"appId" form:"appId" gorm:"column:app_id;comment:应用标识;size:32;"`
	Identify       string    `json:"identify" form:"identify" gorm:"column:identify;comment:模板标识;size:32;"`
	SceneDesc      string    `json:"scene_desc" form:"scene_desc" gorm:"column:scene_desc;comment:场景说明;size:60;"`
	Title          string    `json:"title" form:"title" gorm:"column:title;comment:模板标题;size:120;"`
	Content        string    `json:"content" form:"content" gorm:"column:content;comment:模板消息内容;size:255;"`
	Channel        string    `json:"channel" form:"channel" gorm:"column:channel;comment:通知渠道 wechat:微信 feishu:飞书 getui:个推 web:站内信;size:32;"`
	WechatTempId   string    `json:"wechatTempId" form:"wechatTempId" gorm:"column:wechat_temp_id;comment:微信模板id;size:120;"`
	FeishuUrl      string    `json:"feishuUrl" form:"feishuUrl" gorm:"column:feishu_url;comment:飞书通知地址;size:255;"`
	TargetPlatform string    `json:"targetPlatform" form:"targetPlatform" gorm:"column:target_platform;comment:目标平台 android,ios;size:80;"`
	Status         int       `json:"status" form:"status" gorm:"column:status;comment:1.启用 0.禁用;"`
	Creator        int       `json:"creator" form:"creator" gorm:"column:creator;comment:创建者;size:10;"`
	Updater        int       `json:"updater" form:"updater" gorm:"column:updater;comment:更新者;size:10;"`
	CreatedAt      time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt      time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}
