package nexus

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type Behavior struct {
}

func NewBehavior(config *chg.Config) *Behavior {
	config.InitConfig()
	return &Behavior{}
}

type BehaviorData struct {
	EventName      string `json:"event_name"`
	PageUrl        string `json:"page_url"`
	LaunchScene    int    `json:"launch_scene"`
	UserId         int    `json:"user_id"`
	CardId         int    `json:"card_id"`
	AppName        string `json:"app_name"`
	AppVersion     string `json:"app_version"`
	AppVersionCode string `json:"app_version_code"`
	PlatformType   string `json:"platform_type"`
	DeviceType     string `json:"device_type"`
	DeviceId       string `json:"device_id"`
	OsType         string `json:"os_type"`
	Channel        string `json:"channel"`
	ExternalId     string `json:"external_id"`
	CreatedAt      int64  `json:"created_at"` //毫秒
	EntryAt        int64  `json:"entry_at"`   //毫秒
	LeaveAt        int64  `json:"leave_at"`   //毫秒
}

func (t *Behavior) CreateBehavior(behavior BehaviorData) error {
	mapData := make(map[string]interface{})
	jsonData, err := json.Marshal(behavior)
	if err != nil {
		return err
	}
	mapData["act"] = string(jsonData)
	req := http.Request{
		Method: "POST",
		URL:    "/api/user/behavior",
		Body:   mapData,
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
