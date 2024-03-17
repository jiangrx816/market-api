package service

import (
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

type FormatData struct {
	Value     string `json:"value"`
	Name      string `json:"name"`
	Checked   bool   `json:"checked"`
	Number    int    `json:"number"`
	NumberExt int    `json:"number_ext"`
}

//ApiGetPayList 获取会员价格的列表信息
func (ins *IndexService) ApiGetPayList() (payListData []FormatData) {
	var payList []model.ZMPay
	odb := global.GVA_DB.Model(&model.ZMPay{}).Debug()
	odb = odb.Where("type=1").Order("id asc").Limit(3)
	odb.Find(&payList)

	var temp FormatData
	for idx, _ := range payList {
		temp.Checked = false
		if payList[idx].Id == 1 {
			temp.Checked = true
		}
		temp.Number = payList[idx].Number
		temp.NumberExt = payList[idx].NumberExt
		temp.Value = strconv.Itoa(payList[idx].Id)
		temp.Name = payList[idx].Name + "（￥" + payList[idx].CPrice + "）"
		payListData = append(payListData, temp)
	}
	return payListData
}
