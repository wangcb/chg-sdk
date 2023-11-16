package request

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func SignGenerate(secretKey, method, uri string, body []byte, timestamp int64) string {
	path := uri
	if strings.Contains(uri, "?") {
		parts := strings.Split(uri, "?")
		path = parts[0]
	}

	// Generate the string to sign.
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%d", method, path, string(body), timestamp)
	// Generate the signature.
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}
