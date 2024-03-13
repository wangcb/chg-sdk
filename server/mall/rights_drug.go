package mall

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type RightsDrug struct {
}

func NewRightsDrug(config *chg.Config) *RightsDrug {
	config.InitConfig()
	return &RightsDrug{}
}

// GetRightsDrug 获取权益卡药品清单的药品
func (r *RightsDrug) GetRightsDrug(params map[string]interface{}) (data response.RightsDrugResponse, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/rights/drugs",
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return data, err
	}
	if res.Code != 200 {
		return data, errors.New(res.Msg)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
