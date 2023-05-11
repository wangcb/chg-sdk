package chg

var Configure *Config

type Config struct {
	CoreUrl    string `mapstructure:"core_url" json:"core_url" yaml:"core_url"`          // 用户中心请求地址
	PublicKey  string `mapstructure:"public_key" json:"public_key" yaml:"public_key"`    // jwt解密公钥
	SignSecret string `mapstructure:"sign_secret" json:"sign_secret" yaml:"sign_secret"` // 签名秘钥

	NexusUrl string `mapstructure:"nexus_url" json:"nexus_url" yaml:"nexus_url"` // 服务中心请求地址

	Token string `mapstructure:"token" json:"token"` // jwt
}

func (t *Config) InitConfig() {
	Configure = t
}
