package mall

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type RightsDrug struct {
}

func NewRightsDrug(config *chg.Config) *RightsDrug {
	config.InitConfig()
	return &RightsDrug{}
}

// GetRightsDrug 获取权益卡药品清单的药品
func (r *RightsDrug) GetRightsDrug(params map[string]interface{}) ([]interface{}, error) {
	body := map[string]interface{}{}
	if _, ok := params["card_id"]; ok {
		body["card_id"] = params["card_id"]
	}
	if _, ok := params["keyword"]; ok {
		body["keyword"] = params["keyword"]
	}
	if _, ok := params["medicine_id"]; ok {
		body["medicine_id"] = params["medicine_id"]
	}
	if _, ok := params["medicine_ids"]; ok {
		body["medicine_ids"] = params["medicine_ids"]
	}
	req := http.Request{
		Method: "GET",
		URL:    "/internal/rights-drug-goods",
		Body:   body,
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}
	var data []interface{}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
