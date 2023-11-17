package request

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func SignGenerate(secretKey, method, uri string, body []byte, timestamp int64) string {
	// Generate the string to sign.
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%d", method, uri, string(body), timestamp)
	// Generate the signature.
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}
