package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//该文件为英语跟读绘本的api

//ApiEnglishBookList 英语绘本列表信息
func ApiEnglishBookList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)

	var service service.EnglishService
	list, total := service.GetEnglishBookList(level, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

//ApiEnglishBookInfo 英语绘本详细信息
func ApiEnglishBookInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.EnglishService
	bookInfo := service.GetEnglishBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiOpenId 获取open_id
func ApiOpenId(c *gin.Context) {
	code := c.Query("code")
	var service service.EnglishService
	openId := service.GetOpenId(code)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": openId,
	}, global.SUCCESS_MSG, c)
}