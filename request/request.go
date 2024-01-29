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
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Do(req http.Request, coreUrl string) (*Result, error) {
	timestamp := time.Now().Unix()
	method := strings.ToUpper(req.Method)
	var bodyBytes []byte
	/*if req.Body != nil {
		// 将 map 中的值全部转换为字符串
		if method == "GET" {
			for key, value := range req.Body {
				switch v := value.(type) {
				case int:
					req.Body[key] = strconv.Itoa(v)
				case float64:
					req.Body[key] = strconv.FormatFloat(v, 'f', -1, 64)
				default:
					req.Body[key] = fmt.Sprintf("%v", v)
				}
			}
		}
		bodyBytes, _ = json.Marshal(req.Body)
	}*/
	if method == "GET" || method == "DELETE" {
		parse, err := url.Parse(req.URL)
		if err != nil {
			return nil, err
		}
		params := parse.Query()
		if len(params) > 0 {
			paramMap := make(map[string]string)
			for k, v := range params {
				paramMap[k] = v[0]
			}
			bodyBytes, _ = json.Marshal(paramMap)
		}
	} else {
		bodyBytes, _ = json.Marshal(req.Body)
	}
	sign := SignGenerate(chg.Configure.SignSecret, method, req.URL, bodyBytes, timestamp)
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
	req.URL = coreUrl + req.URL
	client := http.NewClient(30 * time.Second)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	statusCode := resp.StatusCode

	var data Result
	err = json.Unmarshal(resp.Body, &data)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 && statusCode != 201 {
		errMsg := data.Message
		if errMsg == "" {
			errMsg = data.Msg
		}
		return nil, fmt.Errorf("%s", errMsg)
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
