package service

import (
	"fmt"
	"market/global"
	"market/utils"
)

type WechatService struct {
}

func (ws *WechatService) ApiGetWechatData(code string) (wxInfo string) {
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", global.AppId, global.AppSecret, code)
	if request, err := utils.SendGetRequest(urlFormat); err == nil {
		wxInfo = string(request)
	}
	return
}
