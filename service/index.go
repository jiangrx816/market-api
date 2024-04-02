package service

import (
	"context"
	"fmt"
	"market/common/request"
	"market/common/response"
	"market/model"
	"market/utils"
	"strconv"
	"time"
)
import "market/global"

type IndexService struct {
}

//ApiGetCheckLogin 根据openId获取用户是否登录
func (ins *IndexService) ApiGetCheckLogin(openId string) (userInfo model.ZMUser) {
	db := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	db = db.Where("status=1 AND open_id = ?", openId).First(&userInfo)
	userInfo.NickName = utils.TruncateString(userInfo.NickName, 5)
	if userInfo.IsBest == 1 {
		//当前时间
		currentYMD, _ := strconv.Atoi(utils.GetCurrentDateYMD())
		if currentYMD > userInfo.BestLimit {
			userInfo.IsBest = 0
		}
	}
	return
}

//ApiGetUserExtInfo 根据userId获取用户附属信息
func (ins *IndexService) ApiGetUserExtInfo(userId int) (userInfo model.ZMUserExt) {
	db := global.GVA_DB.Model(&model.ZMUserExt{}).Debug()
	db = db.Where("user_id = ?", userId).First(&userInfo)
	return
}

//ApiGetBannerList 获取Banner的列表信息
func (ins *IndexService) ApiGetBannerList(tType int) (bannerList []model.ZMBanner) {
	bookDB := global.GVA_DB.Model(&model.ZMBanner{}).Debug()
	bookDB = bookDB.Where("is_deleted = 0 and status=1 and type = ?", tType).Order("id asc")
	bookDB.Find(&bannerList)
	return bannerList
}

//ApiGetTagList 获取工种的列表信息
func (ins *IndexService) ApiGetTagList() (tagList []model.ZMTags) {
	odb := global.GVA_DB.Model(&model.ZMTags{}).Debug()
	odb = odb.Where("is_deleted = 0 and status=1").Order("sort desc").Order("id asc").Limit(20)
	odb.Find(&tagList)
	return tagList
}

//ApiGetPayList 获取会员价格的列表信息
func (ins *IndexService) ApiGetPayList() (payListData []response.FormatData) {
	var payList []model.ZMPay
	odb := global.GVA_DB.Model(&model.ZMPay{}).Debug()
	odb = odb.Where("type=1 and status=1 and is_deleted = 0").Order("sort desc").Limit(6)
	odb.Find(&payList)

	var temp response.FormatData
	for idx, _ := range payList {
		var payCPrice float64
		payCPrice = payList[idx].CPrice / 100
		temp.Checked = payList[idx].Checked
		temp.TotalFee = payCPrice
		temp.Number = payList[idx].Number
		temp.NumberExt = payList[idx].NumberExt
		temp.Value = strconv.Itoa(payList[idx].Id)
		strValue := fmt.Sprintf("%.2f", payCPrice)
		temp.Name = payList[idx].Name + "（￥" + strValue + "）"
		payListData = append(payListData, temp)
	}
	return payListData
}

//ApiGetGoodPay 获取优选工匠的价格
func (ins *IndexService) ApiGetGoodPay() (info model.ZMPay) {
	var payList []model.ZMPay
	odb := global.GVA_DB.Model(&model.ZMPay{}).Debug()
	odb = odb.Where("type=2 and status=1").Order("sort desc").Limit(1)
	odb.Find(&payList)
	if len(payList) > 0 {
		info = payList[0]
		var payCPrice float64
		payCPrice = info.CPrice / 100
		info.CPrice = payCPrice
	}
	return
}

//ApiGetGoodMemberList 获取优选工匠列表
func (ins *IndexService) ApiGetGoodMemberList(page, tType int) (memberLists []response.MemberData, count int64) {
	tagDataList := ins.GetTagList()
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	var memberList []model.ZMUser
	odb := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	odb = odb.Where("is_deleted = 0 and status= 1 and is_best = 1")
	if tType > 0 && tType < 15 {
		odb = odb.Where(" tag_id = ?", tType)
	}
	odb.Count(&count)
	odb = odb.Order("id desc").Limit(size).Offset(offset)
	odb.Find(&memberList)

	//组合userId的集合
	var userIds []int64
	for idx, _ := range memberList {
		userIds = append(userIds, memberList[idx].UserId)
	}
	var memberExtList []model.ZMUserExt
	odbExt := global.GVA_DB.Model(&model.ZMUserExt{}).Debug()
	odb = odbExt.Where("user_id in(?)", userIds).Find(&memberExtList)

	var temp response.MemberData
	for idx, _ := range memberList {
		temp.Desc = ""
		temp.ViewCount = 1
		temp.Id = memberList[idx].Id
		temp.UserId = memberList[idx].UserId
		temp.OpenId = memberList[idx].OpenId
		temp.NickName = utils.TruncateString(memberList[idx].NickName, 3)
		temp.RealName = memberList[idx].RealName
		temp.HeadUrl = memberList[idx].HeadUrl
		temp.Mobile = memberList[idx].Mobile
		temp.TagId = memberList[idx].TagId
		for dIndex, _ := range tagDataList {
			if memberList[idx].TagId == tagDataList[dIndex].Id {
				temp.TagName = tagDataList[dIndex].Name
			}
		}

		for dIndex, _ := range memberExtList {
			if memberList[idx].UserId == memberExtList[dIndex].UserId {
				temp.Desc = utils.TruncateString(utils.ClearMobileText(utils.RegContent(memberExtList[dIndex].Desc, ins.getBadWordsList())), 50) + "......"
				temp.ViewCount = memberExtList[dIndex].ViewCount
			}
		}
		temp.IsBest = memberList[idx].IsBest
		temp.IsMember = memberList[idx].IsMember
		temp.MemberLimit = memberList[idx].MemberLimit
		memberLists = append(memberLists, temp)
	}

	return
}

//ApiGetMemberInfo 获取会员详情
func (ins *IndexService) ApiGetMemberInfo(userId int) (userInfo response.MemberData) {
	var user model.ZMUser
	odb := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	odb.Where("user_id=?", userId).First(&user)

	var userExt model.ZMUserExt
	odbExt := global.GVA_DB.Model(&model.ZMUserExt{}).Debug()
	odbExt.Where("user_id=?", userId).First(&userExt)

	global.GVA_DB.Model(&model.ZMUserExt{}).Debug().Where("user_id=?", userId).Update("view_count", userExt.ViewCount+1)
	tagInfo := ins.GetTagInfo(user.TagId)
	userInfo.Id = user.Id
	userInfo.UserId = user.UserId
	userInfo.OpenId = user.OpenId
	userInfo.NickName = utils.TruncateString(user.NickName, 5)
	userInfo.RealName = utils.TruncateString(user.RealName, 5)
	userInfo.HeadUrl = user.HeadUrl
	userInfo.Mobile = user.Mobile
	userInfo.TagId = user.TagId
	userInfo.TagName = tagInfo.Name
	userInfo.Desc = utils.ClearMobileText(utils.RegContent(userExt.Desc, ins.getBadWordsList()))
	userInfo.Demo = userExt.Demo
	userInfo.ViewCount = userExt.ViewCount

	return
}

//ApiGetTaskList 获取任务列表
func (ins *IndexService) ApiGetTaskList(page, tType int) (taskLists []response.FormatTaskData, count int64) {
	tagDataList := ins.GetTagList()
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	var taskList []model.ZMTask
	odb := global.GVA_DB.Model(&model.ZMTask{}).Debug()
	odb = odb.Where("is_deleted = 0 and status > 0")
	if tType > 0 && tType < 15 {
		odb = odb.Where(" tag_id = ?", tType)
	}
	odb.Count(&count)
	odb = odb.Order("id desc").Limit(size).Offset(offset)
	odb.Find(&taskList)

	//组合userId的集合
	var userIds []int64
	for idx, _ := range taskList {
		userIds = append(userIds, taskList[idx].UserId)
	}
	userIdsResult := utils.RemoveDuplicates(userIds)
	var memberList []model.ZMUser
	odbUser := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	odb = odbUser.Where("user_id in(?)", userIdsResult).Find(&memberList)

	var temp response.FormatTaskData
	for idx, _ := range taskList {
		temp.Id = taskList[idx].Id
		temp.TagId = taskList[idx].TagId
		temp.UserId = taskList[idx].UserId
		for dIndex, _ := range tagDataList {
			if taskList[idx].TagId == tagDataList[dIndex].Id {
				temp.TagName = tagDataList[dIndex].Name
			}
		}

		//过滤敏感数据并清除手机号
		tempDesc := utils.ClearMobileText(utils.RegContent(taskList[idx].Desc, ins.getBadWordsList()))
		temp.Desc = utils.TruncateString(tempDesc, 60) + "......"

		temp.Mobile = ""
		temp.IsBest = 0
		for dIndex, _ := range memberList {
			if memberList[dIndex].UserId == taskList[idx].UserId {
				temp.Mobile = memberList[dIndex].Mobile
				temp.IsBest = memberList[dIndex].IsBest
			}
		}

		temp.Date = utils.GetUnixTimeToDateTime1(taskList[idx].AddTime)
		temp.Address = taskList[idx].Address
		temp.Status = taskList[idx].Status

		taskLists = append(taskLists, temp)
	}

	return
}

//ApiGetMyTaskList 获取已发布的任务列表
func (ins *IndexService) ApiGetMyTaskList(page, userId int) (taskLists []response.FormatTaskData, count int64) {
	tagDataList := ins.GetTagList()
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	var taskList []model.ZMTask
	odb := global.GVA_DB.Model(&model.ZMTask{}).Debug()
	if userId > 0 {
		odb = odb.Where("is_deleted = 0 and user_id = ?", userId)
	}
	odb.Count(&count)
	odb = odb.Order("id desc").Limit(size).Offset(offset)
	odb.Find(&taskList)

	//组合userId的集合
	var userIds []int64
	for idx, _ := range taskList {
		userIds = append(userIds, taskList[idx].UserId)
	}
	var memberList []model.ZMUser
	odbUser := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	odb = odbUser.Where("user_id in(?)", userIds).Find(&memberList)

	var temp response.FormatTaskData
	for idx, _ := range taskList {
		temp.Id = taskList[idx].Id
		temp.TagId = taskList[idx].TagId
		for dIndex, _ := range tagDataList {
			if taskList[idx].TagId == tagDataList[dIndex].Id {
				temp.TagName = tagDataList[dIndex].Name
			}
		}

		temp.Desc = utils.TruncateString(taskList[idx].Desc, 50) + "......"
		for dIndex, _ := range memberList {
			if taskList[idx].UserId == memberList[dIndex].UserId {
				temp.Mobile = memberList[dIndex].Mobile
			}
		}

		temp.Date = utils.GetUnixTimeToDateTime1(taskList[idx].AddTime)
		temp.Address = utils.TruncateString(taskList[idx].Address, 5)
		temp.Status = taskList[idx].Status
		taskLists = append(taskLists, temp)
	}

	return
}

//ApiGetTaskInfo 获取任务详情
func (ins *IndexService) ApiGetTaskInfo(taskId int) (taskInfo response.FormatTaskData) {
	var task model.ZMTask
	odb := global.GVA_DB.Model(&model.ZMTask{}).Debug()
	odb.Where("id=?", taskId).First(&task)

	var user model.ZMUser
	odbUser := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	odbUser.Where("user_id=?", task.UserId).First(&user)

	tagInfo := ins.GetTagInfo(task.TagId)
	taskInfo.Id = task.Id
	taskInfo.UserId = task.UserId
	taskInfo.TagId = task.TagId
	taskInfo.TagName = tagInfo.Name
	taskInfo.Desc = utils.ClearMobileText(utils.RegContent(task.Desc, ins.getBadWordsList()))
	taskInfo.Mobile = user.Mobile
	taskInfo.Date = utils.GetUnixTimeToDateTime(task.AddTime)
	taskInfo.Address = task.Address
	taskInfo.Status = task.Status

	return
}

//GetTagList 获取所有的工种
func (ins *IndexService) GetTagList() (tagList []model.ZMTags) {
	odb := global.GVA_DB.Model(&model.ZMTags{}).Debug()
	odb.Find(&tagList)
	return tagList
}

//GetTagInfo 获取指定的工种
func (ins *IndexService) GetTagInfo(tagId int) (tagInfo model.ZMTags) {
	odb := global.GVA_DB.Model(&model.ZMTags{}).Debug()
	odb.Where("id=?", tagId).First(&tagInfo)
	return
}

//ApiCheckPushTask 校验是否可发布
func (ins *IndexService) ApiCheckPushTask(userId int) (result bool) {
	s, _ := global.GVA_REDIS.Get(context.Background(), fmt.Sprintf("userPushTask_%d", userId)).Result()

	atoi, _ := strconv.Atoi(s)

	if atoi < 1 {
		result = true
	}

	return
}

//getBadWordsList 获取过滤的坏词
func (ins *IndexService) getBadWordsList() (badWordsSlice []string) {
	var badWords []model.ZMBadWords
	global.GVA_DB.Model(&model.ZMBadWords{}).Debug().Find(&badWords)
	for idx, _ := range badWords {
		badWordsSlice = append(badWordsSlice, badWords[idx].Name)
	}
	return
}

//ApiDoMakeTaskData 发布任务
func (ins *IndexService) ApiDoMakeTaskData(taskData request.MakeTaskData) (result bool) {
	if taskData.Title == "" || taskData.TaskDesc == "" || taskData.Address == "" || taskData.TagId == 0 || taskData.UserId == 0 {
		return
	}

	//判断用户发的次数，5分钟设置最大3次
	cacheCount := fmt.Sprintf("request_count:%d", taskData.UserId)
	count, _ := global.GVA_REDIS.Get(context.Background(), cacheCount).Result()
	countInt, _ := strconv.Atoi(count)
	fmt.Printf("countInt的次数：%#v \n", countInt)
	if countInt >= 3 {
		return false
	}

	var task model.ZMTask
	odb := global.GVA_DB.Model(&model.ZMTask{}).Debug()
	task.TagId = taskData.TagId
	task.UserId = taskData.UserId
	task.Title = taskData.Title
	task.Desc = taskData.TaskDesc
	task.Address = taskData.Address
	task.Status = 1
	task.AddTime = utils.GetCurrentUnixTimestamp()
	affected := odb.Create(&task).RowsAffected
	if affected > 0 {
		result = true
		//键值增加1
		global.GVA_REDIS.Incr(context.Background(), cacheCount).Result()
		//设置5分钟的有效期
		global.GVA_REDIS.Expire(context.Background(), cacheCount, time.Duration(300)*time.Second)
		//单独的防止连续点击
		global.GVA_REDIS.SetNX(context.Background(), fmt.Sprintf("userPushTask_%d", task.UserId), 1, time.Duration(3)*time.Second)
	}
	return
}

//ApiUpdateTaskStatus 更新任务状态
func (ins *IndexService) ApiUpdateTaskStatus(taskData request.UpdateTaskStatus) (result bool) {
	if taskData.TaskId < 0 {
		return
	}
	var task model.ZMTask
	task.Status = taskData.Status
	global.GVA_DB.Model(&model.ZMTask{}).Debug().Where("id=?", taskData.TaskId).Update(&task)
	return true
}

//ApiUpdateMemberData 更新用户资料信息
func (ins *IndexService) ApiUpdateMemberData(memberData request.MemberUpdateData) (result bool) {
	if memberData.UserId <= 0 {
		return
	}
	var member model.ZMUser
	if memberData.NickName != "" {
		member.NickName = memberData.NickName
	}
	if memberData.Mobile != "" {
		member.Mobile = memberData.Mobile
	}
	if memberData.HeadUrl != "" {
		member.HeadUrl = memberData.HeadUrl
	}

	global.GVA_DB.Model(&model.ZMUser{}).Debug().Where("user_id=?", memberData.UserId).Update(&member)
	return true
}

//ApiDoMakeUserData 创建用户
func (ins *IndexService) ApiDoMakeUserData(userData request.MakeUserData) (result bool) {
	if userData.Type < 0 || userData.Mobile == "" || userData.OpenId == "" {
		return
	}
	//判断用户的openID是否存在
	var userTemp model.ZMUser
	global.GVA_DB.Model(&model.ZMUser{}).Debug().Where("open_id=?", userData.OpenId).First(&userTemp)
	if userTemp.UserId > 0 {
		return true
	}

	var user model.ZMUser
	odb := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	var userIdTemp = utils.GetCurrentUnixTimestamp()
	user.UserId = userIdTemp
	user.OpenId = userData.OpenId
	if userData.Type == 2 {
		user.OpenId = "app-" + strconv.FormatInt(userIdTemp, 10)
	}
	user.Mobile = userData.Mobile
	user.Type = userData.Type
	if userData.NickName != "" && len(userData.NickName) > 1 {
		user.NickName = userData.NickName
	}
	if userData.HeadImg != "" && len(userData.HeadImg) > 10 {
		user.HeadUrl = userData.HeadImg
	}
	affected := odb.Create(&user).RowsAffected
	if affected > 0 {
		result = true
	}
	return
}
