package config

type Wechat struct {
	AppId   string `mapstructure:"appid" json:"appid" yaml:"appid"`          // appid
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 端口值
	MchId   string `mapstructure:"mch-id" json:"mch-id" yaml:"mch-id"`       // 商户号
	MchCert string `mapstructure:"mch-cert" json:"mch-cert" yaml:"mch-cert"` // 序列号
	MchIv3  string `mapstructure:"mch-iv3" json:"mch-iv3" yaml:"mch-iv3"`    // iv4
}
