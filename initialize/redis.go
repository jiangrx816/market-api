package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"wechat/global"
	"time"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
		PoolSize: 100,
	})
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := RedisClient.Ping(cxt).Result()
	if err != nil {
		log.Fatal("redis ping err::", err)
	}
	log.Printf("redis init on port :%+v", redisCfg.Port)
	global.GVA_REDIS = RedisClient
}
