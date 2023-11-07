package chg

var Configure *Config

type Config struct {
	CoreUrl    string `mapstructure:"core_url" json:"core_url" yaml:"core_url"`          // 用户中心请求地址
	NexusUrl   string `mapstructure:"nexus_url" json:"nexus_url" yaml:"nexus_url"`       // 服务中心请求地址
	MallUrl    string `mapstructure:"mall_url" json:"mall_url" yaml:"mall_url"`          // 商城
	JwtSecret  string `mapstructure:"jwt_secret" json:"jwt_secret" yaml:"jwt_secret"`    // jwt解密公钥
	SignSecret string `mapstructure:"sign_secret" json:"sign_secret" yaml:"sign_secret"` // 签名秘钥
	Token      string `mapstructure:"token" json:"token"`                                // jwt
}

func (t *Config) InitConfig() {
	Configure = t
}
