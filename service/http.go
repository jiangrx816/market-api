package service

import (
	"encoding/json"
	"strconv"
	"wechat/global"
	"wechat/model"
	"wechat/utils"
)

type HttpService struct {
}

type HttpPostResult struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data map[int]Item `json:"data"`
}

type Item struct {
	ID           int    `json:"id"`
	BookID       int    `json:"book_id"`
	CNContent    string `json:"cn_content"`
	ENContent    string `json:"en_content"`
	SpellContent string `json:"spell_content"`
	CNArr        string `json:"cn_arr"`
	ENArr        string `json:"en_arr"`
	PlayTime     int    `json:"play_time"`
	ENPlayTime   int    `json:"en_play_time"`
	Img          string `json:"img"`
	Meaning      string `json:"meaning"`
	MeaningCNArr string `json:"meaning_cn_arr"`
	MeaningEN    string `json:"meaning_en"`
	MeaningENArr string `json:"meaning_en_arr"`
	Rank         int    `json:"rank"`
	IsShow       int    `json:"is_show"`
}



//SendHttpPost 发送post请求
func (hs *HttpService) SendHttpPost() {
	var poetryList []model.PoetryPicture
	global.GVA_DB.Model(&model.PoetryPicture{}).Debug().Find(&poetryList)

	for idx, _ := range poetryList {
		poetryId := strconv.Itoa(poetryList[idx].BookId)
		hs.DealHttpPostData(poetryId)
	}
}

// DealHttpPostData 处理http返回的数据
func (hs *HttpService) DealHttpPostData(bookId string) {
	url := "https://mzbook.com/api/book.Book/getSentence"
	var paramList = []utils.HttpParams{
		{
			Param: "id",
			Item:  bookId,
		},
	}
	postJson := utils.SendHttpPost(url, paramList)
	var res HttpPostResult
	json.Unmarshal(postJson, &res)

	listData := res.Data
	var poetryPicture model.PoetryPictureInfo
	var poetryPictureList []model.PoetryPictureInfo
	for idx, _ := range listData {
		poetryPicture.BookId = listData[idx].BookID
		poetryPicture.CN = listData[idx].CNContent
		poetryPicture.Pic = listData[idx].Img
		poetryPicture.Mp3 = listData[idx].CNArr
		poetryPicture.Position = listData[idx].Rank
		poetryPictureList = append(poetryPictureList, poetryPicture)
	}
	db := global.GVA_DB.Model(&model.PoetryPictureInfo{}).Debug()
	for idx, _ := range poetryPictureList {
		db.Create(&poetryPictureList[idx])
	}
}
