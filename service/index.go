package service

import (
	"wechat/global"
	"wechat/model"
)

type IndexService struct {

}

func (is *IndexService)ApiIndex(categoryId int) (bookList []model.PoetryPicture, total int64) {
	bookDB := global.GVA_DB.Model(&model.PoetryPicture{}).Debug()
	bookDB = bookDB.Where("type_id = ?", 1).Count(&total)
	bookDB = bookDB.Limit(1).Offset(10)
	bookDB.Find(&bookList)
	return
}