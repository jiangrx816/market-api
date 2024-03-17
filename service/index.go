package service

import "market/model"
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
func (ins *IndexService) ApiGetPayList() (payList []model.ZMPay) {
	odb := global.GVA_DB.Model(&model.ZMPay{}).Debug()
	odb = odb.Where("type=1").Order("id asc").Limit(3)
	odb.Find(&payList)
	return payList
}
