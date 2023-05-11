package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    map[string]interface{}
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

type Client struct {
	HTTPClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) Do(req Request) (*Response, error) {
	httpRequest, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		httpRequest.Header.Set(key, value)
	}

	if req.Body != nil {
		contentType := req.Headers["Content-Type"]
		if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			formData := url.Values{}
			for key, value := range req.Body {
				formData.Set(key, fmt.Sprintf("%v", value))
			}
			httpRequest.Body = io.NopCloser(strings.NewReader(formData.Encode()))
		} else {
			body, err := json.Marshal(req.Body)
			if err != nil {
				return nil, err
			}
			httpRequest.Body = io.NopCloser(bytes.NewBuffer(body))
			httpRequest.Header.Set("Content-Type", "application/json")
		}
		defer httpRequest.Body.Close()
	}

	httpResponse, err := c.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	bodyBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	for key, values := range httpResponse.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &Response{
		StatusCode: httpResponse.StatusCode,
		Headers:    headers,
		Body:       bodyBytes,
	}, nil
}
