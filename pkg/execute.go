package pkg

import (
	"github.com/robfig/cron"
	"log"
	"strconv"
	"wechat/global"
	"wechat/initialize"
	"wechat/router"
)

func Execute() {
	//初始化加载配置信息
	initialize.InitViperConfig()
	//初始化连接数据库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		defer global.GVA_DB.DB().Close()
	}
	//初始化连接Redis
	if global.GVA_CONFIG.SYSTEM.UseRedis == true {
		initialize.Redis()
		//是否使用路由缓存
		if global.GVA_CONFIG.SYSTEM.UseHttpCache == true {
			initialize.HttpCache()
		}
	}
	//初始化路由
	r := router.InitRouter()
	//启动WEB服务
	if err := r.Run(":" + strconv.Itoa(global.GVA_CONFIG.SYSTEM.Port)); err != nil {
		log.Fatal("服务器启动失败...Error:" + err.Error())
	}
}

func ExecuteCron() {
	// 创建一个 cron 实例
	c := cron.New()
	//// 添加定时任务
	//var service service.CacheService
	//
	////每30分钟执行一次
	//c.AddFunc("@every 5m", func() {
	//	fmt.Println("Cron job executed at:", time.Now())
	//	service.DealRedisRouterCache("TimeLong")
	//})
	// 启动 cron
	c.Start()
	//在程序退出时关闭 cron
	defer c.Stop()
}
