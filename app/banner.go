package app

import (
	"github.com/gin-gonic/gin"
	"market/common"
	"market/global"
	"market/service"
)

//ApiGetBannerList 获取Banner列表
func ApiGetBannerList(c *gin.Context) {
	var service service.BannerService
	list := service.ApiGetBannerList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}
