package common

import (
	"github.com/gin-gonic/gin"
	"wechat/global"
	"strings"
)

//CheckRequestUserAgent 检查是否微信来源
func CheckRequestUserAgent(c *gin.Context) bool {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), global.MINI_WECHAT)
	if !isFlag {
		ReturnResponse(global.FORBID, map[string]interface{}{}, global.FORBID_MSG, c)
		return false
	}
	return true
}
