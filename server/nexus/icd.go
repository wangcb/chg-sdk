package nexus

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type Icd struct {
}

func NewIcd(config *chg.Config) *Icd {
	config.InitConfig()
	return &Icd{}
}

// GetIcd 查询疾病
func (t *Icd) GetIcd(params map[string]string) (icdList []response.Icd, err error) {

	req := http.Request{
		Method: "GET",
		URL:    request.BuildURL("/api/getIcd", params),
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &icdList)
	if err != nil {
		return nil, err
	}

	return icdList, nil
}
