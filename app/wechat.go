package app

import (
	"github.com/gin-gonic/gin"
	"market/common"
	"market/common/request"
	"market/global"
	"market/service"
	"net/http"
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
	var json request.MakePhotoData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var service service.WechatService

	data := service.ApiGetWxUserPhoneNumber(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}
