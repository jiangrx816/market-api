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
		api.GET("/poetry/school/getOpenId", app.ApiSchoolOpenId)
		api.GET("/poetry/school/getPoetryLog", app.ApiPoetryLog)
		api.POST("/poetry/school/uploadMp3", app.ApiUploadPoetryMp3)
		api.POST("/poetry/school/addVideoLog", app.ApiAddPoetryVideoLog)
		//初高中
		api.GET("/poetry/junior/getList", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryList)
		api.GET("/poetry/junior/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryInfo)
		//成语
		api.GET("/poetry/cheng/getList", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryList)
		api.GET("/poetry/cheng/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryInfo)

		//英文绘本
		api.GET("/english/getList", app.ApiEnglishBookList)
		api.GET("/english/getBookInfo", app.ApiEnglishBookInfo)
		api.GET("/english/getOpenId", app.ApiOpenId)

		//下载图片
		api.GET("/downLoad/pic", app.ApiDownLoadPic)
		api.GET("/http/post", app.ApiHttpPost)
	}

	//路由组v2 校验是否微信或者小程序请求访问
	apiV2 := router.Group("v2")
	apiV2.Use(middleware.CheckWechatMiddleware())
	{
		//古诗词成语
		//小学
		apiV2.GET("/poetry/school/getList", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryList)
		apiV2.GET("/poetry/school/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryInfo)
		apiV2.GET("/poetry/school/getOpenId", app.ApiSchoolOpenId)
		apiV2.GET("/poetry/school/getPoetryLog", app.ApiPoetryLog)
		apiV2.POST("/poetry/school/uploadMp3", app.ApiUploadPoetryMp3)
		apiV2.POST("/poetry/school/addVideoLog", app.ApiAddPoetryVideoLog)
		//初高中
		apiV2.GET("/poetry/junior/getList", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryList)
		apiV2.GET("/poetry/junior/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryInfo)
		//成语
		apiV2.GET("/poetry/cheng/getList", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryList)
		apiV2.GET("/poetry/cheng/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryInfo)

		//中文绘本
		apiV2.GET("/chinese/getNavList", app.ApiChineseNavList)
		apiV2.GET("/chinese/getList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookList)
		apiV2.GET("/chinese/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookInfo)
		//中文绘本专辑
		apiV2.GET("/chinese/getAlbumList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumList)
		apiV2.GET("/chinese/getAlbumListInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumListInfo)
		apiV2.GET("/chinese/getAlbumInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumInfo)
		//古诗绘本
		apiV2.GET("/poetry/book/getList", routerCache(global.RedisURL_CACHE), app.ApiPoetryBookList)
		apiV2.GET("/poetry/book/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiPoetryBookInfo)

		//英文绘本
		apiV2.GET("/english/getList", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookList)
		apiV2.GET("/english/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookInfo)
		apiV2.GET("/english/getOpenId", app.ApiOpenId)

		//百科知识
		apiV2.GET("/baike/getCategoryCount", routerCache(global.RedisURL_CACHE), app.ApiCategoryCount)
		apiV2.GET("/baike/getQuestion", app.ApiQuestion)
	}
	return router
}
