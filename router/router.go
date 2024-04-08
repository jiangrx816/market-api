package router

import (
	"github.com/gin-gonic/gin"
	"market/app"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.MaxMultipartMemory = 32 << 20 // 32MB
	//路由组v1
	api := router.Group("api/v1")

	api.Use()
	{
		//获取地址列表
		api.GET("/getAddressHot", app.ApiGetAddressHot)
		api.GET("/getAddressList", app.ApiGetAddressList)
		api.GET("/getAddressChild", app.ApiGetAddressChildList)
		//根据openId获取用户是否登录
		api.GET("/getCheckLogin", app.ApiGetCheckLogin)
		//根据userId获取用户附属信息
		api.GET("/getUserExt", app.ApiGetUserExtInfo)
		//上传文件
		api.POST("/uploadFile", app.ApiUploadFileData)
		//获取banner列表
		api.GET("/getBannerList", app.ApiGetBannerList)
		api.GET("/getBannerListNew", app.ApiGetBannerListNew)
		//获取工种列表
		api.GET("/getTagList", app.ApiGetTagList)
		api.GET("/getTagSelect", app.ApiGetTagSelect)
		//获取会员价格列表
		api.GET("/getPayList", app.ApiGetPayList)
		//获取优选工匠的价格
		api.GET("/getGoodPay", app.ApiGetGoodPay)
		//获取优选工匠列表
		api.GET("/getMemberList", app.ApiGetGoodMemberList)
		//获取会员详情
		api.GET("/getMemberInfo", app.ApiGetMemberInfo)
		//更新用户资料信息
		api.POST("/updateMemberData", app.ApiUpdateMemberData)
		//获取任务列表
		api.GET("/getTaskList", app.ApiGetTaskList)
		//获取已发布的任务列表
		api.GET("/getMyTaskList", app.ApiGetMyTaskList)
		//更新任务状态
		api.POST("/updateTaskStatus", app.ApiUpdateTaskStatus)
		//获取任务详情
		api.GET("/getTaskInfo", app.ApiGetTaskInfo)
		//发布任务
		api.POST("/doMakeTaskData", app.ApiDoMakeTaskData)
		//代发任务
		api.POST("/other/doMakeTaskData", app.ApiDoMakeOtherTaskData)
		//校验是否可发布
		api.GET("/checkPushTask", app.ApiCheckPushTask)
		//创建用户
		api.POST("/doMakeUserData", app.ApiDoMakeUserData)
	}
	//路由组v1
	apiWxApi := router.Group("api/wechat")
	apiWxApi.Use()
	{
		//获取openid
		apiWxApi.GET("/getWxData", app.ApiGetWechatData)
		//获取access_token
		apiWxApi.GET("/getWxAccessToken", app.ApiGetWxAccessToken)
		//获取用户手机号
		apiWxApi.POST("/getWxUserPhoneNumber", app.ApiGetWxUserPhoneNumber)
		//微信支付
		apiWxApi.POST("/pay", app.ApiGetWxPay)
		//微信开通优选工匠
		apiWxApi.POST("/open/pay", app.ApiGetWxOpenPay)
		//微信支付回调
		apiWxApi.POST("/pay/notice", app.ApiGetWxPayCallback)
		//微信支付更新为取消
		apiWxApi.POST("/pay/cancel", app.ApiGetWxPayCancel)
		//微信支付更新为取消
		apiWxApi.POST("/pay/refunds", app.ApiGetWxPayRefunds)
	}

	return router
}
