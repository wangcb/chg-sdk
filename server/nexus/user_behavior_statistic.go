package nexus

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type UserBehaviorStatistic struct{}

func NewUserBehaviorStatistic(config *chg.Config) *UserBehaviorStatistic {
	return &UserBehaviorStatistic{}
}

// GetUserBehaviorStatisticList 查询用户行为统计
func (t *UserBehaviorStatistic) GetUserBehaviorStatisticList(params request.UserBehaviorStatisticSearch) (list []response.UserBehaviorStatistic, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "user/behaviorStatistic",
		Body:   StructToMap(params),
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// StructToMap 结构体转map[string]interface{}
func StructToMap(obj interface{}) map[string]interface{} {
	b, _ := json.Marshal(obj)
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	return m
}
