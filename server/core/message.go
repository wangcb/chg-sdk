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
func (t *Message) TemplateList(params map[string]any) (list any, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/message/template",
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	return res.Data, nil
}

// TemplateAdd 新增模版
func (t *Message) TemplateAdd(params map[string]any) error {
	req := http.Request{
		Method: "POST",
		URL:    "/api/message/template",
		Body:   params,
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
func (t *Message) TemplateEdit(params map[string]any) error {
	req := http.Request{
		Method: "PUT",
		URL:    "/api/message/template",
		Body:   params,
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
func (t *Message) TemplateDetail(id int) (data any, err error) {
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
	return res.Data, nil
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

// Send 消息发送
func (t *Message) Send(template string, data map[string]interface{}, toUser interface{}) error {
	// 将 data 转换为 JSON 字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := map[string]interface{}{
		"identify": template,
		"data":     string(jsonData),
	}
	switch v := toUser.(type) {
	case []int:
		body["to_users"] = v
	case int:
		body["to_user"] = v
	}
	req := http.Request{
		Method: "POST",
		URL:    "/api/message/send",
		Body:   body,
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
func (t *Message) ReadMessage(params map[string]any) error {
	req := http.Request{
		Method: "POST",
		URL:    "/api/message/read",
		Body:   params,
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

// MessageList 获取消息列表
func (t *Message) MessageList(params map[string]any) (list any, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/message",
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	return res.Data, nil
}

// MessageDetail 消息详情
func (t *Message) MessageDetail(id int) (data any, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/message/" + strconv.Itoa(id),
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return data, err
	}
	if res.Code != 200 {
		return data, errors.New(res.Message)
	}
	return res.Data, nil
}
