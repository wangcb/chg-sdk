package mall

import (
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"strconv"
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
			"user_card_id": strconv.Itoa(int(userCardId)),
			"user_id":      strconv.Itoa(int(userId)),
			"rights_no":    rightsNo,
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
