package config

type Wechat struct {
	AppId        string `mapstructure:"appid" json:"appid" yaml:"appid"`                         // appid
	Secret       string `mapstructure:"secret" json:"secret" yaml:"secret"`                      // 端口值
	SchoolAppId  string `mapstructure:"school_appid" json:"school_appid" yaml:"school_appid"`    // appid
	SchoolSecret string `mapstructure:"school_secret" json:"school_secret" yaml:"school_secret"` // 端口值
}
