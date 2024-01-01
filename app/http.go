package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/global"
	"wechat/service"
)

//ApiHttpPost 发起http post请求
func ApiHttpPost(c *gin.Context) {
	var service service.HttpService
	service.SendHttpPost()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{}, global.SUCCESS_MSG, c)
}
