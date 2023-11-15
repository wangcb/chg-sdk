package core

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type Applet struct {
}

func NewApplet(config *chg.Config) *Applet {
	config.InitConfig()
	return &Applet{}
}

type SubscribeMessage struct {
	ToUser           string                 `json:"to_user"`
	TemplateID       string                 `json:"template_id"`
	Page             string                 `json:"page"`
	MiniProgramState string                 `json:"mini_program_state"`
	Data             map[string]interface{} `json:"data"`
}

// SendSubscribeMessage 发送小程序订阅消息
func (t *Applet) SendSubscribeMessage(data SubscribeMessage, platform string) (resp response.ResponseSendSucscribeMessage, err error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/mina/send-subscribe-msg",
		Body: map[string]interface{}{
			"to_user":            data.ToUser,
			"template_id":        data.TemplateID,
			"page":               data.Page,
			"mini_program_state": data.MiniProgramState,
			"data":               data.Data,
		},
		Headers: map[string]string{
			"Platform": platform,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return resp, err
	}
	if res.Code != 200 {
		return resp, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
