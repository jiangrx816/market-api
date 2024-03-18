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
		//获取优选工匠列表
		api.GET("/getMemberList", app.ApiGetGoodMemberList)
		//获取会员详情
		api.GET("/getMemberInfo", app.ApiGetMemberInfo)
		//获取任务列表
		api.GET("/getTaskList", app.ApiGetTaskList)
		//获取任务详情
		api.GET("/getTaskInfo", app.ApiGetTaskInfo)
		//发布任务
		api.POST("/doMakeTaskData", app.ApiDoMakeTaskData)
		//校验是否可发布
		api.GET("/checkPushTask", app.ApiCheckPushTask)
	}

	return router
}
