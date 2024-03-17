package router

import (
	"github.com/gin-gonic/gin"
	"market/app"
	"market/global"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	//路由组v1
	api := router.Group("api/v1")
	api.Use()
	{
		//上传文件
		api.POST("/uploadFile", app.ApiUploadFileData)
		//获取banner列表
		api.GET("/getBannerList", routerCache(global.RedisURL_CACHE), app.ApiGetBannerList)
		//获取工种列表
		api.GET("/getTagList", routerCache(global.RedisURL_CACHE), app.ApiGetTagList)
		//获取会员价格列表
		api.GET("/getPayList", app.ApiGetPayList)
	}

	return router
}
