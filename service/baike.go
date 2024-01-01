package service

import (
	"context"
	"fmt"
	"strconv"
	"wechat/common/response"
	"wechat/global"
	"wechat/model"
)

type BaiKeService struct {
}

//GetCategoryCount
func (bs *BaiKeService) GetCategoryCount() (categoryData []response.CategoryData) {
	var categoryList []model.Category
	db := global.GVA_DB.Model(&model.Category{}).Debug()
	db.Where("status = ?", 1).Find(&categoryList)

	db1 := global.GVA_DB.Model(&model.BaiKe{}).Debug()
	var categoryDataList []response.CategoryData
	db1.Raw("SELECT category_id, COUNT(id) AS category_count FROM s_baike GROUP BY category_id").Scan(&categoryDataList)

	var temp response.CategoryData
	for _, item := range categoryList {
		temp.CategoryId = strconv.Itoa(item.Id)
		temp.CategoryCount = "0"
		categoryData = append(categoryData, temp)
	}

	for index, item := range categoryData {
		for _, it := range categoryDataList {
			if item.CategoryId == it.CategoryId {
				categoryData[index].CategoryCount = it.CategoryCount
			}
		}
	}

	return
}

// GetLPopData 从对应栏目中的队列中lpop数据
func (bs *BaiKeService) GetLPopData(categoryId int) map[string]interface{} {
	//获取之前判断队列中还有多少数据
	bs.IsCheckCount(categoryId)

	//队列名称
	queue := fmt.Sprintf(global.DEFAULT_QUEUE)
	if categoryId > 0 {
		queue = fmt.Sprintf(global.QUEUE, categoryId)
	}

	questionId, err := global.GVA_REDIS.LPop(context.Background(), queue).Result()
	if err != nil {
		fmt.Println("从队列中获取数据失败")
	}
	// 创建db
	var baiKe model.BaiKe
	db := global.GVA_DB.Model(&model.BaiKe{}).Debug()
	db = db.Where("id = ?", questionId).Find(&baiKe)

	return map[string]interface{}{
		"id":          baiKe.Id,
		"category_id": baiKe.CategoryId,
		"question":    baiKe.Question,
		"option_a":    baiKe.OptionA,
		"option_b":    baiKe.OptionB,
		"option_c":    baiKe.OptionC,
		"option_d":    baiKe.OptionD,
		"answer":      baiKe.Answer,
		"analytic":    baiKe.Analytic,
	}
}

// IsCheckCount 校验队列中的数据是否小于指定的数量
func (bs *BaiKeService) IsCheckCount(categoryId int) {
	//队列名称
	queue := global.DEFAULT_QUEUE
	if categoryId > 0 {
		queue = fmt.Sprintf(global.QUEUE, categoryId)
	}
	total, _ := global.GVA_REDIS.LLen(context.Background(), queue).Result()
	if total < global.QUEUE_LEN {
		bs.PushDataToQueue(categoryId)
	}
}

// PushDataToQueue 将对应栏目中的数据push到队列中
func (bs *BaiKeService) PushDataToQueue(categoryId int) error {
	// 创建db
	var baiKeList []model.BaiKe
	db := global.GVA_DB.Model(&model.BaiKe{}).Debug()
	if categoryId > 0 {
		db = db.Select("id").Where("category_id = ?", categoryId).Order("question desc, id desc").Find(&baiKeList)
	} else {
		db = db.Select("id").Where("category_id != 21").Order("question desc, id desc").Find(&baiKeList)
	}

	questionIds := make([]int, 0)
	for _, item := range baiKeList {
		questionIds = append(questionIds, item.Id)
	}

	//队列名称
	queue := fmt.Sprintf(global.DEFAULT_QUEUE)
	if categoryId > 0 {
		queue = fmt.Sprintf(global.QUEUE, categoryId)
	}

	pipe := global.GVA_REDIS.Pipeline()
	for _, item := range questionIds {
		pipe.RPush(context.Background(), queue, item)
	}
	if _, err := pipe.Exec(context.Background()); err != nil {
		return err
	}

	return nil
}