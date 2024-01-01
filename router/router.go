package router

import (
	"github.com/gin-gonic/gin"
	"market/app"
	"market/global"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	//路由组v1
	api := router.Group("api")
	api.Use()
	{
		//获取banner列表
		api.GET("/getBannerList", routerCache(global.RedisURL_CACHE), app.ApiGetBannerList)
	}

	return router
}
