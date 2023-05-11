package nexus

import (
	"fmt"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
)

type QiNiu struct {
}

func NewQiNiu(config *chg.Config) *QiNiu {
	config.InitConfig()
	return &QiNiu{}
}

func (t *QiNiu) MakeQiniuPrivateUrl(keys map[string]string, bucket string) (map[string]string, error) {
	if len(keys) == 0 {
		return keys, nil
	}
	req := http.Request{
		Method: "POST",
		URL:    "/api/qiniu/private",
		Body: map[string]interface{}{
			"keys": keys,
			"name": bucket,
		},
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}

	stringData := make(map[string]string)
	for key, value := range res.Data.(map[string]interface{}) {
		stringData[key] = fmt.Sprintf("%v", value)
	}
	return stringData, nil
}
