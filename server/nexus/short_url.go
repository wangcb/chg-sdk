package nexus

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type ShortUrl struct {
}

func NewShortUrl(config *chg.Config) *ShortUrl {
	return &ShortUrl{}
}

// CreateShortUrl 创建短链接
func (s *ShortUrl) CreateShortUrl(params map[string]interface{}) error {
	req := http.Request{
		Method: "POST",
		URL:    request.BuildURL("/api/shortUrl", nil),
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// UpdateShortUrl 更新短链接
func (s *ShortUrl) UpdateShortUrl(params map[string]interface{}) error {
	req := http.Request{
		Method: "PUT",
		URL:    request.BuildURL("/api/shortUrl", nil),
		Body:   params,
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// GetShortUrlList 获取短链接列表
func (s *ShortUrl) GetShortUrlList(params map[string]string) (list any, err error) {
	req := http.Request{
		Method: "GET",
		URL:    request.BuildURL("/api/shortUrl", params),
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	list = res.Data

	return
}

// GetShortUrl 获取短链接
func (s *ShortUrl) GetShortUrl(id string) (data map[string]interface{}, err error) {
	req := http.Request{
		Method: "GET",
		URL:    request.BuildURL("/api/shortUrl/"+id, nil),
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	bytes, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return
}

// DeleteShortUrl 删除短链接
func (s *ShortUrl) DeleteShortUrl(param map[string]string) error {
	req := http.Request{
		Method: "DELETE",
		URL:    request.BuildURL("/api/shortUrl", param),
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	return nil
}

// DeleteShortUrls 批量删除短链接
func (s *ShortUrl) DeleteShortUrls(ids []int) error {
	req := http.Request{
		Method: "DELETE",
		URL:    request.BuildURL("/api/delShortUrls", nil),
		Body: map[string]interface{}{
			"ids": ids,
		},
	}
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	return nil
}

// GenerateShortUrl 生成短连接
func (s *ShortUrl) GenerateShortUrl() (string, error) {
	req := http.Request{
		Method: "POST",
		URL:    request.BuildURL("/api/generateShortUrlCode", nil),
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return res.Data.(string), nil
}


