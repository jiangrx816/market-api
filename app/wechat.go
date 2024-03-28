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

//ApiGetWxAccessToken 获取access_token
func ApiGetWxAccessToken(c *gin.Context) {
	var service service.WechatService
	data := service.ApiGetWxAccessToken()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}

//ApiGetWxUserPhoneNumber 获取用户手机号
func ApiGetWxUserPhoneNumber(c *gin.Context) {
	var service service.WechatService
	token, _ := c.GetPostForm("token")
	code, _ := c.GetPostForm("code")
	data := service.ApiGetWxUserPhoneNumber(code, token)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}
