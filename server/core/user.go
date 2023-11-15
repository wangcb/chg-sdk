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

// RegWithPhone 手机号注册
func (t *User) RegWithPhone(phone string, source string, platform string) (token string, err error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/token/phone",
		Body: map[string]interface{}{
			"phone": phone,
		},
		Headers: map[string]string{
			"Source":   source,
			"Platform": platform,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return token, err
	}
	if res.Code != 200 {
		return token, errors.New(res.Message)
	}
	return res.Data.(string), nil
}

// GetUserOpenid 获取用户openid
func (t *User) GetUserOpenid(userId int, appid string) (info response.UserWechatInfo, err error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/user-wechat",
		Body: map[string]interface{}{
			"user_id": userId,
			"app_id":  appid,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return info, err
	}
	if res.Code != 200 {
		return info, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return info, err
	}
	return info, nil
}
