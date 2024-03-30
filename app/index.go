package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"market/common"
	"market/common/request"
	"market/global"
	"market/service"
	"market/utils"
	"net/http"
	"path/filepath"
	"time"
)

//ApiGetCheckLogin 根据openId获取用户是否登录
func ApiGetCheckLogin(c *gin.Context) {
	openId := c.Query("open_id")
	var service service.IndexService
	info := service.ApiGetCheckLogin(openId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": info,
	}, global.SUCCESS_MSG, c)
}

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

//ApiGetGoodMemberList 获取优选工匠列表
func ApiGetGoodMemberList(c *gin.Context) {
	page := utils.GetIntParamItem("page", 10, c)
	tType := utils.GetIntParamItem("type", 0, c)
	var service service.IndexService
	list, count := service.ApiGetGoodMemberList(page, tType)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":  list,
		"count": count,
	}, global.SUCCESS_MSG, c)
}

//ApiGetMemberInfo 获取会员详情
func ApiGetMemberInfo(c *gin.Context) {
	userId := utils.GetIntParamItem("user_id", 0, c)
	var service service.IndexService
	info := service.ApiGetMemberInfo(userId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": info,
	}, global.SUCCESS_MSG, c)
}

//ApiGetTaskList 获取任务列表
func ApiGetTaskList(c *gin.Context) {
	page := utils.GetIntParamItem("page", 10, c)
	tType := utils.GetIntParamItem("type", 0, c)
	var service service.IndexService
	list, count := service.ApiGetTaskList(page, tType)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":  list,
		"count": count,
	}, global.SUCCESS_MSG, c)
}

//ApiGetMyTaskList 获取已发布的任务列表
func ApiGetMyTaskList(c *gin.Context) {
	page := utils.GetIntParamItem("page", 10, c)
	userId := utils.GetIntParamItem("user_id", 0, c)
	var service service.IndexService
	list, count := service.ApiGetMyTaskList(page, userId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":  list,
		"count": count,
	}, global.SUCCESS_MSG, c)
}

//ApiGetTaskInfo 获取任务详情
func ApiGetTaskInfo(c *gin.Context) {
	taskId := utils.GetIntParamItem("task_id", 0, c)
	var service service.IndexService
	info := service.ApiGetTaskInfo(taskId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": info,
	}, global.SUCCESS_MSG, c)
}

//ApiDoMakeTaskData 发布任务
func ApiDoMakeTaskData(c *gin.Context) {
	var json request.MakeTaskData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var service service.IndexService

	log.Println(json)
	res := service.ApiDoMakeTaskData(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"result": res,
	}, global.SUCCESS_MSG, c)
}

//ApiUpdateTaskStatus 更新任务状态
func ApiUpdateTaskStatus(c *gin.Context) {
	var json request.UpdateTaskStatus
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.IndexService
	res := service.ApiUpdateTaskStatus(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"result": res,
	}, global.SUCCESS_MSG, c)
}

//ApiUpdateMemberData 更新用户资料信息
func ApiUpdateMemberData(c *gin.Context) {
	var json request.MemberUpdateData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.IndexService

	res := service.ApiUpdateMemberData(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"result": res,
	}, global.SUCCESS_MSG, c)
}

//ApiCheckPushTask 校验是否可发布
func ApiCheckPushTask(c *gin.Context) {
	userId := utils.GetIntParamItem("user_id", 0, c)
	var service service.IndexService
	info := service.ApiCheckPushTask(userId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": info,
	}, global.SUCCESS_MSG, c)
}

//ApiUploadFileData 上传录音
func ApiUploadFileData(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		common.ReturnResponse(global.FAIL, map[string]interface{}{}, global.FAIL_MSG, c)
		return
	}
	// 获取文件名
	fileName := filepath.Base(file.Filename)
	// 获取文件扩展名
	extension := filepath.Ext(fileName)

	// 生成文件名（使用时间戳）
	fName := time.Now().Format("20060102150405") + extension

	path := "/data/static/market"

	// 保存文件到服务器
	c.SaveUploadedFile(file, filepath.Join(path, fName))

	dst := "https://oss.58haha.com/market/" + fName

	c.JSON(200, gin.H{
		"dst": dst,
	})
}

//ApiDoMakeUserData 创建用户
func ApiDoMakeUserData(c *gin.Context) {
	var json request.MakeUserData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.IndexService
	res := service.ApiDoMakeUserData(json)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"result": res,
	}, global.SUCCESS_MSG, c)
}
