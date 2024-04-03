package app

import (
	"encoding/json"
	"fmt"
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

//ApiGetWxPay 微信支付
func ApiGetWxPay(c *gin.Context) {
	var json request.WXPayData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.WechatService
	data := service.ApiCreateWxPay(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}

//ApiGetWxOpenPay 微信开通优选工匠
func ApiGetWxOpenPay(c *gin.Context) {
	var json request.OpenGoodPay
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.WechatService
	data := service.ApiGetWxOpenPay(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}

//ApiGetWxPayCallback 微信支付通知
func ApiGetWxPayCallback(c *gin.Context) {
	var service service.WechatService
	notifyReq, err := service.ApiDealWxPayCallback(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "FAIL", "message": err.Error()})
		return
	}
	//解析Plaintext参数，这里面可以拿到订单的基本信息
	var data = []byte(notifyReq.Resource.Plaintext)
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("JSON转换失败：", err)
		return
	}
	// 处理通知内容
	fmt.Println("解密结果", notifyReq)
	fmt.Println("解密结果Plaintext", result)

	//将解密结果进行处理
	service.ApiDealUserPaySuccess(notifyReq, result)
	c.JSON(http.StatusOK, gin.H{"code": "SUCCESS"})
}

//ApiGetWxPayCancel 微信支付更新为取消
func ApiGetWxPayCancel(c *gin.Context) {
	var json request.WXCancelPayData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.WechatService
	data := service.ApiGetWxPayCancel(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
	}, global.SUCCESS_MSG, c)
}

//ApiGetWxPayRefunds 微信支付退款
func ApiGetWxPayRefunds(c *gin.Context) {
	var json request.WXRefundsPayData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.WechatService
	message, code := service.ApiGetWxPayRefunds(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"message": message,
		"code":    code,
	}, global.SUCCESS_MSG, c)
}
