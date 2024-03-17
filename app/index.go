package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"market/common"
	"market/global"
	"market/service"
	"path"
	"strings"
)

//ApiGetBannerList 获取Banner列表
func ApiGetBannerList(c *gin.Context) {
	var service service.IndexService
	list := service.ApiGetBannerList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

//ApiGetTagList 获取工种列表
func ApiGetTagList(c *gin.Context) {
	var service service.IndexService
	list := service.ApiGetTagList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

//ApiGetPayList 获取会员价格列表
func ApiGetPayList(c *gin.Context) {
	var service service.IndexService
	list := service.ApiGetPayList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

//ApiUploadFileData 上传录音
func ApiUploadFileData(c *gin.Context) {
	file, err := c.FormFile("file")
	filePath := "market/member/ext"
	if err == nil {
		Path := fmt.Sprintf("/data/static/%s", filePath)
		dst := path.Join(Path, file.Filename)
		fmt.Printf("file.Filename:%s \n", file.Filename)
		fmt.Printf("dst:%s \n", dst)
		c.SaveUploadedFile(file, dst)
		dst = strings.Replace(dst, Path, "https://static.58haha.com/poetry_log", 1)
		fmt.Printf("dst:%s \n", dst)
		c.JSON(200, gin.H{
			"dst": dst,
		})
	} else {
		common.ReturnResponse(global.FAIL, map[string]interface{}{}, global.FAIL_MSG, c)
	}
}
