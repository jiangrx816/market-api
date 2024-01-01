package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wechat/cache"
	"wechat/global"
)

func routerCache(sec int64) (res gin.HandlerFunc) {
	fmt.Printf("sec:%v", sec)
	return cache.CacheByRequestURI(global.GVA_HTTP_CACHE, time.Duration(sec)*time.Hour)
}
