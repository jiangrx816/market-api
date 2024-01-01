package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"wechat/common/response"
	"wechat/global"
	"wechat/model"
)

type EnglishService struct {
}

//GetEnglishBookList 获取英语绘本的列表信息
func (es *EnglishService) GetEnglishBookList(level, page int) (englishBookList []response.ResponseEnglishBook, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	var bookList []model.EnglishBook
	bookDB := global.GVA_DB.Model(&model.EnglishBook{}).Debug()
	bookDB = bookDB.Where("level = ?", level).Count(&total)
	bookDB = bookDB.Order("position desc").Limit(size).Offset(offset)
	bookDB.Find(&bookList)

	var bookInfoCountList []response.ResponseBookInfoCount
	infoDB := global.GVA_DB.Model(&model.ChineseBookInfo{}).Debug()
	infoDB.Raw("SELECT book_id,count(id) as book_count FROM s_huiben_info GROUP BY book_id").Scan(&bookInfoCountList)

	var temp response.ResponseEnglishBook
	for _, item := range bookList {
		temp.Id = item.Id
		temp.BookId = item.BookId
		temp.Title = item.Title
		temp.Icon = item.Icon
		temp.Level = item.Level
		temp.Position = item.Position
		englishBookList = append(englishBookList, temp)
	}

	for index, item := range englishBookList {
		for _, it := range bookInfoCountList {
			if item.BookId == it.BookId {
				englishBookList[index].BookCount = it.BookCount
			}
		}
	}

	sort.Slice(englishBookList, func(i, j int) bool {
		if englishBookList[i].Position > englishBookList[j].Position {
			return true
		}
		return englishBookList[i].Position == englishBookList[j].Position && englishBookList[i].Id < englishBookList[j].Id
	})

	return
}

//GetEnglishBookInfo 获取英语绘本的详情信息
func (es *EnglishService) GetEnglishBookInfo(bookId string) (bookInfoItems []model.EnglishBookInfo) {
	db := global.GVA_DB.Model(&model.EnglishBookInfo{}).Debug()
	db = db.Where("book_id = ?", bookId).Order("id asc").Find(&bookInfoItems)
	return
}

//GetOpenId 获取open_id信息
func (es *EnglishService) GetOpenId(code string) (openId string) {
	var data response.OpenIdData
	appid := global.GVA_CONFIG.Wechat.AppId
	secret := global.GVA_CONFIG.Wechat.Secret
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &data)
	openId = data.Openid
	return
}
