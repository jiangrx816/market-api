package service

import (
	"market/common/response"
	"market/model"
	"strconv"
)
import "market/global"

type IndexService struct {
}

//ApiGetBannerList 获取Banner的列表信息
func (ins *IndexService) ApiGetBannerList() (bannerList []model.ZMBanner) {
	bookDB := global.GVA_DB.Model(&model.ZMBanner{}).Debug()
	bookDB = bookDB.Where("status=1").Order("id asc")
	bookDB.Find(&bannerList)
	return bannerList
}

//ApiGetTagList 获取工种的列表信息
func (ins *IndexService) ApiGetTagList() (tagList []model.ZMTags) {
	odb := global.GVA_DB.Model(&model.ZMTags{}).Debug()
	odb = odb.Where("status=1").Order("id asc").Limit(15)
	odb.Find(&tagList)
	return tagList
}

//ApiGetPayList 获取会员价格的列表信息
func (ins *IndexService) ApiGetPayList() (payListData []response.FormatData) {
	var payList []model.ZMPay
	odb := global.GVA_DB.Model(&model.ZMPay{}).Debug()
	odb = odb.Where("type=1 and status=1").Order("sort desc").Limit(6)
	odb.Find(&payList)

	var temp response.FormatData
	for idx, _ := range payList {
		temp.Checked = payList[idx].Checked
		temp.Number = payList[idx].Number
		temp.NumberExt = payList[idx].NumberExt
		temp.Value = strconv.Itoa(payList[idx].Id)
		temp.Name = payList[idx].Name + "（￥" + payList[idx].CPrice + "）"
		payListData = append(payListData, temp)
	}
	return payListData
}

//ApiGetGoodMemberList 获取优选工匠列表
func (ins *IndexService) ApiGetGoodMemberList(page, tType int) (memberLists []response.MemberData) {

	tagDataList := ins.GetTagList()
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	var memberList []model.ZMUser
	odb := global.GVA_DB.Model(&model.ZMUser{}).Debug()
	odb = odb.Where("status= 1 and is_best = 1")
	if tType > 0 {
		odb = odb.Where(" tag_id = ?", tType)
	}
	odb = odb.Order("id desc").Limit(size).Offset(offset)
	odb.Find(&memberList)

	//组合userId的集合
	var userIds []int
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
		temp.NickName = memberList[idx].NickName
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
				temp.Desc = memberExtList[dIndex].Desc
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
	userInfo.NickName = user.NickName
	userInfo.RealName = user.RealName
	userInfo.HeadUrl = user.HeadUrl
	userInfo.Mobile = user.Mobile
	userInfo.TagId = user.TagId
	userInfo.TagName = tagInfo.Name
	userInfo.Desc = userExt.Desc
	userInfo.Demo = userExt.Demo
	userInfo.ViewCount = userExt.ViewCount

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
