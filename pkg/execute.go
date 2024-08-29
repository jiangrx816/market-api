package pkg

import (
	"log"
	"market/global"
	"market/initialize"
	"market/router"
	"strconv"
)

func Execute() {
	//初始化加载配置信息
	initialize.InitViperConfig()
	//初始化连接数据库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		defer global.GVA_DB.DB().Close()
	}
	////初始化连接Redis
	//if global.GVA_CONFIG.SYSTEM.UseRedis == true {
	//	initialize.Redis()
	//	//是否使用路由缓存
	//	if global.GVA_CONFIG.SYSTEM.UseHttpCache == true {
	//		initialize.HttpCache()
	//	}
	//}
	//初始化路由
	r := router.InitRouter()
	//启动WEB服务
	if err := r.Run(":" + strconv.Itoa(global.GVA_CONFIG.SYSTEM.Port)); err != nil {
		log.Fatal("服务器启动失败...Error:" + err.Error())
	}
}
