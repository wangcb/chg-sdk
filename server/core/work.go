package core

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
	"strconv"
)

type Work struct {
}

func NewWork(config *chg.Config) *Work {
	config.InitConfig()
	return &Work{}
}

// GetWorkUserList 获取企业微信员工列表
func (t *Work) GetWorkUserList(platform string, cursor string, limit int) (users response.WorkUserList, err error) {
	req := http.Request{
		Method: "GET",
		URL:    "/api/work/users",
		Body: map[string]interface{}{
			"cursor": cursor,
			"limit":  strconv.Itoa(limit),
		},
		Headers: map[string]string{
			"Platform": platform,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return users, err
	}
	if res.Code != 200 {
		return users, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return users, err
	}
	if users.ErrMsg != "ok" {
		return users, errors.New("查询失败")
	}
	data := res.Data.(map[string]interface{})
	deptUsers := data["dept_user"].([]interface{})
	for _, item := range deptUsers {
		var user response.WorkUserListUser
		if userId, ok := item.(map[string]interface{})["userid"]; ok {
			user.UserId = userId.(string)
		}
		if department, ok := item.(map[string]interface{})["department"]; ok {
			user.Department = int64(department.(float64))
		}
		users.DeptUser = append(users.DeptUser, user)
	}
	return users, nil
}

// ContactQrcode 企业微信联系人二维码
func (t *Work) ContactQrcode(platform string, userId string, state string) (qrcode response.ContactQrcode, err error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/work/contact/qrcode",
		Body: map[string]interface{}{
			"userid": userId,
			"state":  state,
		},
		Headers: map[string]string{
			"Platform": platform,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return qrcode, err
	}
	if res.Code != 200 {
		return qrcode, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &qrcode)
	if err != nil {
		return qrcode, err
	}
	if qrcode.ErrCode != 0 {
		return qrcode, errors.New(qrcode.ErrMsg)
	}
	return qrcode, nil
}

// TagCustomer 给外部联系人打标签
func (t *Work) TagCustomer(platform string, userId string, externalUserId string, tag []string) (err error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/work/contact/tag-user",
		Body: map[string]interface{}{
			"userid":           userId,
			"tag":              tag,
			"external_user_id": externalUserId,
		},
		Headers: map[string]string{
			"Platform": platform,
		},
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	var tagRes response.WorkCommon
	err = json.Unmarshal(bytes, &tagRes)
	if err != nil {
		return err
	}
	if tagRes.ErrCode != 0 {
		return errors.New(tagRes.ErrMsg)
	}
	return nil
}
