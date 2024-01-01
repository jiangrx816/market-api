package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/global"
	"wechat/service"
)

//ApiCache 模拟缓存
func ApiCache(c *gin.Context) {
	var service service.CacheService
	service.GetCateInfo()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{}, global.SUCCESS_MSG, c)
}
