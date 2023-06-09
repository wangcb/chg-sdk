package nexus

import (
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type Sms struct {
}

func NewSms(config *chg.Config) *Sms {
	config.InitConfig()
	return &Sms{}
}

// Verify 验证短信验证码
func (t *Sms) Verify(phone string, code string) error {
	req := http.Request{
		Method: "POST",
		URL:    "/api/sms/code/verify",
		Body: map[string]interface{}{
			"phone": phone,
			"code":  code,
		},
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// Send 发送短信
func (t *Sms) Send(phone string, template string, params map[string]interface{}, args ...string) error {
	signName := ""
	if len(args) > 0 {
		signName = args[0]
	}
	req := http.Request{
		Method: "POST",
		URL:    "/api/sms/send",
		Body: map[string]interface{}{
			"phone":     phone,
			"template":  template,
			"data":      params,
			"sign_name": signName,
		},
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}
