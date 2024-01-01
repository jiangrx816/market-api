package middleware

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
)

//CheckWechatMiddleware 验证是否为微信小程序访问
func CheckWechatMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !common.CheckRequestUserAgent(ctx) {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
