package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangcb/chg-sdk/request"
	"io"
	"net/http"
	"strconv"
	"time"
)

func CheckSign(c *gin.Context, secretKey string, expireTime int) error {
	// 获取签名和时间戳
	signature := c.GetHeader("Sign")
	timestampStr := c.GetHeader("Timestamp")

	// 如果签名或时间戳不存在，则直接返回错误
	if signature == "" || timestampStr == "" {
		return errors.New("signature or timestamp missing")
	}
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return errors.New("invalid timestamp")
	}
	// 如果签名已过期，则返回错误
	expiration := time.Duration(expireTime) * time.Second

	// 检查签名是否过期
	if time.Now().Sub(time.Unix(timestamp, 0)) > expiration {
		return errors.New("signature has expired")
	}
	// 获取参数
	var jsonBytes []byte
	if c.Request.Body == http.NoBody {
		// 处理 GET 请求参数
		params := c.Request.URL.Query()
		paramMap := make(map[string]string)
		for k, v := range params {
			paramMap[k] = v[0]
		}
		jsonBytes, err = json.Marshal(paramMap)
		if err != nil {
			return err
		}
	} else {
		// 读取请求体内容
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return errors.New("failed to read request body")
		}
		// 可以自动ASCII码排序
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if len(bodyBytes) > 0 {
			// 解析请求体内容
			var jsonParams map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &jsonParams); err != nil {
				return errors.New("invalid request body")
			}
			jsonBytes, err = json.Marshal(jsonParams)
			if err != nil {
				return err
			}
		}
	}

	calcSig := request.SignGenerate(secretKey, c.Request.Method, c.Request.URL.Path, jsonBytes, timestamp)
	if calcSig != signature {
		return errors.New("invalid signature")
	}
	return nil
}
