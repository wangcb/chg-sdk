package robot

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type RobotCall struct {
}

func NewRobotCall(config *chg.Config) *RobotCall {
	config.InitConfig()
	return &RobotCall{}
}

func (r *RobotCall) MobileCall(params request.RobotCall) (callId string, err error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/robot/call",
		Body: map[string]interface{}{
			"phone":      params.Phone,
			"voice_code": params.VoiceCode,
			"source":     params.Source,
			"remark":     params.Remark,
		},
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	var data response.RobotCall
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	return data.CallId, nil
}
