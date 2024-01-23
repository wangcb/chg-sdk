package mall

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type RightsCard struct {
}

func NewRightsCard(config *chg.Config) *RightsCard {
	config.InitConfig()
	return &RightsCard{}
}

// HasUserRights 获取用户权益
func (r *RightsCard) HasUserRights(userCardId int64, userId int64, rightsNo string) (bool, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/user/rights",
		Body: map[string]interface{}{
			"user_card_id": userCardId,
			"user_id":      userId,
			"rights_no":    rightsNo,
		},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return false, err
	}
	return res.Data.(bool), nil
}

// RightsUsedCallback 权益使用回调
func (r *RightsCard) RightsUsedCallback(userCardId int64, rightsNo string, userId int64, amount int) (bool, error) {
	req := http.Request{
		Method: "POST",
		URL:    "/internal/user/rights/notify",
		Body: map[string]interface{}{
			"user_card_id": userCardId,
			"rights_no":    rightsNo,
			"user_id":      userId,
			"amount":       amount,
		},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return false, err
	}
	if res.Code != 200 {
		return false, errors.New(res.Msg)
	}
	return true, nil
}

// RightsUserCardList 用户权益卡列表
func (r *RightsCard) RightsUserCardList(params map[string]interface{}) (data *response.UserRightsCardList, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/user/rights/list",
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}

	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return
}

// RightsUserCardDetails 用户权益卡详情
func (r *RightsCard) RightsUserCardDetails(id uint) (data *response.UserRightsCard, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/user/rights/details",
		Body:   map[string]interface{}{"id": id},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}

	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return
}

// RightsChannelList 权益渠道列表
func (r *RightsCard) RightsChannelList(params map[string]interface{}) (data []response.RightsChannel, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/rights/channel",
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}

	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return
}
