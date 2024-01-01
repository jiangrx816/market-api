package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//ApiDownLoadPic 下载图片
func ApiDownLoadPic(c *gin.Context) {
	var service service.DownLoadService
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	service.GetDownLoadImages(page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{}, global.SUCCESS_MSG, c)
}