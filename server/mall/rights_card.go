package mall

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
	"strconv"
)

type RightsCard struct {
}

func NewRightsCard(config *chg.Config) *RightsCard {
	config.InitConfig()
	return &RightsCard{}
}

// GetUserRights 获取用户权益
func (r *RightsCard) GetUserRights(userCardId int64, userId int64) (*response.UserRightsCard, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/rights",
		Body: map[string]interface{}{
			"user_card_id": strconv.Itoa(int(userCardId)),
			"user_id":      strconv.Itoa(int(userId)),
		},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}
	var data response.UserRightsCard
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// RightsUsedCallback 权益使用回调
func (r *RightsCard) RightsUsedCallback(userCardId int64, rightsNo string, userId int64, amount int) (*response.UserRights, error) {
	req := http.Request{
		Method: "POST",
		URL:    "/internal/rights/notify",
		Body: map[string]interface{}{
			"user_card_id": userCardId,
			"rights_no":    rightsNo,
			"user_id":      userId,
			"amount":       amount,
		},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}
	var data response.UserRights
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
