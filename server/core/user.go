package core

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type User struct {
}

func NewUser(config *chg.Config) *User {
	config.InitConfig()
	return &User{}
}

// GetInfo 获取用户信息
func (t *User) GetInfo() (*response.UserInfo, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/user",
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var data response.UserInfo
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetRealName 获取实名信息
func (t *User) GetRealName() (*response.RealName, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/real-name",
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var data response.RealName
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
