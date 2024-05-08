package server

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type Common struct {
}

func NewCommon(config *chg.Config) *Common {
	config.InitConfig()
	return &Common{}
}

func (s *Common) Request(url string, params map[string]any, method string) (res any, err error) {
	req := http.Request{
		Method: method,
		URL:    url,
		Body:   params,
	}
	res, err = request.Do(req, chg.Configure.NexusUrl)
	return
}

func (s *Common) NexusRequest(url string, params map[string]any, method string) (data map[string]interface{}, err error) {
	req := http.Request{
		Method: method,
		URL:    url,
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
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
	return data, nil
}

func (s *Common) NovaRequest(url string, params map[string]any, method string) (data map[string]interface{}, err error) {
	req := http.Request{
		Method: method,
		URL:    url,
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.NovaUrl)
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
	return data, nil
}

func (s *Common) CloudRequest(url string, params map[string]any, method string) (data map[string]interface{}, err error) {
	req := http.Request{
		Method: method,
		URL:    url,
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.CloudUrl)
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
	return data, nil
}

func (s *Common) CoreUserRequest(url string, params map[string]any, method string, header map[string]string) (data map[string]interface{}, err error) {
	req := http.Request{
		Method: method,
		URL:    url,
		Body:   params,
	}
	if header != nil {
		req.Headers = header
	}
	res, err := request.Do(req, chg.Configure.CoreUrl)
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
	return data, nil
}
