package nexus

import (
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type ShortUrl struct {
}

func NewShortUrl(config *chg.Config) *ShortUrl {
	config.InitConfig()
	return &ShortUrl{}
}

func (s *ShortUrl) GeneralRequest(url string, params map[string]any, method string) (list any, err error) {
	req := http.Request{
		Method: method,
		URL:    url,
		Body:   params,
	}
	list, err = request.Do(req, chg.Configure.NexusUrl)
	return
}
