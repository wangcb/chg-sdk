package core

import (
	"errors"
	jwtpkg "github.com/golang-jwt/jwt/v4"
	"github.com/wangcb/chg-sdk/chg"
	"os"
	"strings"
	"time"
)

type Auth struct {
}

func NewAuth(config *chg.Config) *Auth {
	config.InitConfig()
	return &Auth{}
}

type JWTCustomClaims struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"user_name"`
	Extra    string `json:"extra"`
	jwtpkg.RegisteredClaims
}

func (t *Auth) ParseToken(token string) (*JWTCustomClaims, error) {
	if token == "" {
		return nil, errors.New("需要认证才能访问")
	}
	type JWT struct {
		Secret interface{}
		// 刷新 Token 的最大过期时间
		MaxRefresh time.Duration
	}
	jwt := &JWT{}
	secret := chg.Configure.JwtSecret
	if strings.HasSuffix(secret, ".pem") {
		publicKeyFile, err := os.ReadFile(secret)
		if err != nil {
			return nil, errors.New("获取公钥文件失败")
		}
		publicKey, err := jwtpkg.ParseRSAPublicKeyFromPEM(publicKeyFile)
		if err != nil {
			return nil, errors.New("解析公钥文件失败")
		}
		jwt.Secret = publicKey
	} else {
		jwt.Secret = []byte(secret)
	}
	// 按空格分割
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("请求头中 Authorization 格式有误")
	}
	jwtToken, err := jwtpkg.ParseWithClaims(parts[1], &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.Secret, nil
	})
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, errors.New("请求令牌格式有误")
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, errors.New("令牌已过期")
			}
		}
		return nil, errors.New("请求令牌无效")
	}
	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := jwtToken.Claims.(*JWTCustomClaims); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, errors.New("请求令牌无效")
}
