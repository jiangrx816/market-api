package router

import (
	"github.com/gin-gonic/gin"
	"wechat/app"
	"wechat/global"
	"wechat/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	//路由组v1
	api := router.Group("v1")
	api.Use()
	{
		//首页demo
		api.GET("/index", app.ApiIndex)
		api.GET("/index/cache",routerCache(300), app.ApiIndex)
		api.GET("/index/cache1",routerCache(600), app.ApiIndex)

		//古诗词成语
		//小学
		api.GET("/poetry/school/getList", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryList)
		api.GET("/poetry/school/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryInfo)
	}


	return router
}
