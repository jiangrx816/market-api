package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//该文件为百科知识的api

//ApiCategoryCount 栏目下所对应的数据量
func ApiCategoryCount(c *gin.Context) {
	var service service.BaiKeService
	list := service.GetCategoryCount()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": list,
	}, global.SUCCESS_MSG, c)
}

// ApiQuestion 获取对应的栏目答题数据
func ApiQuestion(c *gin.Context) {
	categoryId := utils.GetIntParamItem("category_id", global.DEFAULT_BOOK_ID, c)
	var service service.BaiKeService
	question := service.GetLPopData(categoryId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": question,
	}, global.SUCCESS_MSG, c)
}