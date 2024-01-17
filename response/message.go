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
	ID             int64     `json:"id" form:"id"`
	AppId          string    `json:"appId" form:"appId"`
	Identify       string    `json:"identify" form:"identify" `
	SceneDesc      string    `json:"scene_desc" form:"scene_desc"`
	Title          string    `json:"title" form:"title"`
	Content        string    `json:"content" form:"content"`
	Channel        string    `json:"channel" form:"channel"`
	WechatTempId   string    `json:"wechatTempId" form:"wechatTempId"`
	FeishuUrl      string    `json:"feishuUrl" form:"feishuUrl"`
	TargetPlatform string    `json:"targetPlatform" form:"targetPlatform"`
	Status         int       `json:"status" form:"status"`
	Creator        int       `json:"creator" form:"creator"`
	Updater        int       `json:"updater" form:"updater"`
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
