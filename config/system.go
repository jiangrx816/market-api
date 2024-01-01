package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境值
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                               // 端口值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	UseHttpCache bool   `mapstructure:"use-http-cache" json:"use-http-cache" yaml:"use-http-cache"` // 使用use-http-cache
}