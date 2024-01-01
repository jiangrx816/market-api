package global

import (
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"wechat/config"
)

var (
	GVA_CONFIG     *config.Server
	GVA_DB         *gorm.DB
	GVA_REDIS      *redis.Client
	GVA_LOG        *zap.Logger
	GVA_HTTP_CACHE *persist.RedisStore
)
