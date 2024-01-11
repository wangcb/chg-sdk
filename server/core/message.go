package core

import (
	"chg-sdk/request"
	"core-user/app/model/message"
	messageRes "core-user/app/request"
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"strconv"
)

type Message struct {
}

func NewMessage(config *chg.Config) *Message {
	config.InitConfig()
	return &Message{}
}

// TemplateList 获取模板配置列表
func (t *Message) TemplateList(params messageRes.TemplateList) (*message.MessageTemplate, error) {
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
	var data message.MessageTemplate
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// TemplateAdd 模板配置新增
func (t *Message) TemplateAdd(params messageRes.TemplateParam) error {
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
func (t *Message) TemplateEdit(params messageRes.TemplateParam) error {
	req := http.Request{
		Method: "PUT",
		URL:    "/api/message/template",
		Body: map[string]interface{}{
			"id":              params.Id,
			"appid":           params.Appid,
			"identify":        params.Identify,
			"scene_desc":      params.SceneDesc,
			"channel":         params.Channel,
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
