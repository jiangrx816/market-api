package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"path"
	"strings"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//该文件为古诗词成语的api

//ApiSchoolPoetryList 小学古诗列表信息
func ApiSchoolPoetryList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	var service service.PoetryService
	list, total := service.GetSchoolPoetryList(page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

//ApiSchoolPoetryInfo 小学古诗详细信息
func ApiSchoolPoetryInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_BOOK_ID, c)
	var service service.PoetryService
	bookInfo := service.GetSchoolPoetryInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiJuniorPoetryList 中学古诗列表信息
func ApiJuniorPoetryList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)

	var service service.PoetryService
	list, total := service.GetJuniorPoetryList(page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

//ApiJuniorPoetryInfo 中学古诗详细信息
func ApiJuniorPoetryInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_BOOK_ID, c)
	var service service.PoetryService
	bookInfo := service.GetJuniorPoetryInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiChengPoetryList 成语列表信息
func ApiChengPoetryList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)

	var service service.PoetryService
	list, total := service.GetChengPoetryList(level, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

//ApiChengPoetryInfo 成语详细信息
func ApiChengPoetryInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_BOOK_ID, c)
	var service service.PoetryService
	bookInfo := service.ChengPoetryInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiSchoolOpenId 获取appid
func ApiSchoolOpenId(c *gin.Context) {
	code := c.Query("code")
	var service service.PoetryService
	openId := service.GetSchoolOpenId(code)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": openId,
	}, global.SUCCESS_MSG, c)
}

//ApiPoetryLog
func ApiPoetryLog(c *gin.Context) {
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info":  "",
		"total": 1,
	}, global.SUCCESS_MSG, c)
}

//ApiUploadPoetryMp3 上传录音
func ApiUploadPoetryMp3(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		var Path string = "/data/web/static/poetry_log"
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

//ApiAddPoetryVideoLog
func ApiAddPoetryVideoLog(c *gin.Context) {
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{}, global.SUCCESS_MSG, c)
}