package initialize

import (
	"github.com/chenyahui/gin-cache/persist"
	"log"
	"wechat/global"
)

func HttpCache() {
	//判断redis是连接成功
	if global.GVA_REDIS == nil {
		log.Printf("%+v", "redis server not connect, http-cache failed.")
		return
	}
	log.Printf("%+v", "http-cache started.")
	global.GVA_HTTP_CACHE = persist.NewRedisStore(global.GVA_REDIS)
}
