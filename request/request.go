package request

import (
	"encoding/json"
	"fmt"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
	Data    interface{}
}

func Do(req http.Request, url string) (*Result, error) {
	timestamp := time.Now().Unix()
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = json.Marshal(req.Body)
	}
	sign := SignGenerate(chg.Configure.SignSecret, req.Method, req.URL, bodyBytes, timestamp)
	// 组装header 信息
	headers := make(map[string]string)
	if req.Headers != nil {
		headers = req.Headers
	}
	headers["Content-Type"] = "application/json"
	headers["Sign"] = sign
	headers["Timestamp"] = strconv.Itoa(int(timestamp))
	if chg.Configure.Token != "" {
		headers["Authorization"] = chg.Configure.Token
	}
	req.Headers = headers

	// 请求地址拼接
	req.URL = url + req.URL
	client := http.NewClient(30 * time.Second)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var data Result
	err = json.Unmarshal(resp.Body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func BuildURL(urlPath string, params map[string]string) string {
	var paramStrings []string

	// 遍历map，将键值对拼接成"key=value"形式的字符串
	for key, value := range params {
		paramStrings = append(paramStrings, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	}

	return fmt.Sprintf("%s?%s", urlPath, strings.Join(paramStrings, "&"))
}
