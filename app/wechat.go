package app

import (
	"github.com/gin-gonic/gin"
	"market/common"
	"market/global"
	"market/service"
)

//ApiGetWechatData 根据code换取 openId, sessionKey, unionId
func ApiGetWechatData(c *gin.Context) {
	code := c.Query("code")
	var service service.WechatService
	data := service.ApiGetWechatData(code)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}
