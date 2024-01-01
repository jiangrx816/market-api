package service

import "market/model"
import "market/global"

type BannerService struct {
}

//ApiGetBannerList 获取Banner的列表信息
func (cs *BannerService) ApiGetBannerList() (bannerList []model.ZMBanner) {
	bookDB := global.GVA_DB.Model(&model.ZMBanner{}).Debug()
	bookDB = bookDB.Where("status=1").Order("id asc")
	bookDB.Find(&bannerList)
	return bannerList
}
