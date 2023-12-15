package nexus

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type UserBehaviorStatistic struct{}

func NewUserBehaviorStatistic(config *chg.Config) *UserBehaviorStatistic {
	config.InitConfig()
	return &UserBehaviorStatistic{}
}

// GetUserBehaviorStatisticList 查询用户行为统计
func (t *UserBehaviorStatistic) GetUserBehaviorStatisticList(params map[string]interface{}) (list map[string]interface{}, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/user/behaviorStatistic",
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	bytes, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bytes, &list); err != nil {
		return nil, err
	}

	return list, nil
}
