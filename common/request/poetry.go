package request

import "wechat/model"

type PoetryVideoReq struct {
	OpenId   string `json:"open_id" comment:"open_id"`
	PoetryId int    `json:"poetry_id" comment:"poetry_id"`
	Mp3      string `json:"mp3" comment:"mp3"`
}

func (req PoetryVideoReq) GeneratePoetryVideoLog(model *model.PoetryLog) {
	model.OpenId = req.OpenId
	model.PoetryId = req.PoetryId
	model.Mp3 = req.Mp3
}
