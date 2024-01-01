package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wechat/common"
	"wechat/global"
	"wechat/service"
)

func ApiIndex(c *gin.Context) {
	//检查是否微信请求来源
	//if !common.CheckRequestUserAgent(c) {
	//	return
	//}
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	var service service.IndexService
	index, total := service.ApiIndex(categoryId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":  index,
		"total": total,
	}, global.SUCCESS_MSG, c)
}
