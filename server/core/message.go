package core

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
	"strconv"
)

type Message struct {
}

func NewMessage(config *chg.Config) *Message {
	config.InitConfig()
	return &Message{}
}

// GetWechatTemplate 根据应用获取微信后台配置的模板列表 （微信官方配置的模板）
func (t *Message) GetWechatTemplate(platForm string) ([]response.WechatTemplate, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/message/getWechatTemplate",
		Headers: map[string]string{
			"Platform": platForm,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var data []response.WechatTemplate
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// TemplateList 获取模板配置列表
func (t *Message) TemplateList(params request.TemplateList) (data []*response.MessageTemplate, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/message/template",
		Body: map[string]interface{}{
			"page":    params.Page,
			"size":    params.PageSize,
			"appid":   params.Appid,
			"channel": params.Channel,
			"status":  params.Status,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// TemplateAdd 模板配置新增
func (t *Message) TemplateAdd(params request.TemplateParam) error {
	req := http.Request{
		Method: "POST",
		URL:    "/api/message/template",
		Body: map[string]interface{}{
			"appid":           params.Appid,
			"identify":        params.Identify,
			"scene_desc":      params.SceneDesc,
			"channel":         params.Channel,
			"title":           params.Title,
			"content":         params.Content,
			"temp_params":     params.TempParams,
			"wechat_temp_id":  params.WechatTempId,
			"feishu_url":      params.FeishuUrl,
			"target_platform": params.TargetPlatform,
			"status":          params.Status,
			"creator":         params.Creator,
			"updater":         params.Updater,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// TemplateEdit 模板配置修改
func (t *Message) TemplateEdit(params request.TemplateParam) error {
	req := http.Request{
		Method: "PUT",
		URL:    "/api/message/template",
		Body: map[string]interface{}{
			"id":              params.Id,
			"appid":           params.Appid,
			"identify":        params.Identify,
			"scene_desc":      params.SceneDesc,
			"channel":         params.Channel,
			"temp_params":     params.TempParams,
			"title":           params.Title,
			"content":         params.Content,
			"wechat_temp_id":  params.WechatTempId,
			"feishu_url":      params.FeishuUrl,
			"target_platform": params.TargetPlatform,
			"status":          params.Status,
			"updater":         params.Updater,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// TemplateDetail 模板详情
func (t *Message) TemplateDetail(id int) (data *response.MessageTemplate, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/message/template/" + strconv.Itoa(id),
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// TemplateDel 模板删除
func (t *Message) TemplateDel(id int) error {
	req := http.Request{
		Method: "DELETE",
		URL:    "/api/message/template/" + strconv.Itoa(id),
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// SendMessage 消息发送
// 参数示例
//
//	{
//		"identify": "test", 必传
//		"push_type": 1,
//		"to_user": "[\"999\"]", push_type = 3 可不传
//		"request_param": "[{\"thing2\":\"每日一次\"},{\"thing4\":\"您可以在微信中直接回复医生\"}]"  thing2 根据模板中的temp_params 参数进行拼装
//	}
func (t *Message) SendMessage(params request.SendMessage) error {
	req := http.Request{
		Method: "POST",
		URL:    "/api/message/send",
		Body: map[string]interface{}{
			"identify":      params.Identify,
			"push_type":     params.PushType,
			"to_user":       params.ToUser,
			"request_param": params.RequestParam,
			"content":       params.Content,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// ReadMessage 消息已读
func (t *Message) ReadMessage(params request.ReadUser) error {
	req := http.Request{
		Method: "POST",
		URL:    "/api/message/read",
		Body: map[string]interface{}{
			"user_id":    params.UserId,
			"message_id": params.MessageId,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}
