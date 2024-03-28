package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"market/global"
	"market/utils"
	"net/http"
)

type WechatService struct {
}

//ApiGetWechatData 根据code换取 openId, sessionKey, unionId
func (ws *WechatService) ApiGetWechatData(code string) (wxInfo string) {
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", global.AppId, global.AppSecret, code)
	if request, err := utils.SendGetRequest(urlFormat); err == nil {
		wxInfo = string(request)
	}
	return
}

//ApiGetWxAccessToken 获取access_token
func (ws *WechatService) ApiGetWxAccessToken() (wxInfo string) {
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", global.AppId, global.AppSecret)
	if request, err := utils.SendGetRequest(urlFormat); err == nil {
		wxInfo = string(request)
	}
	return
}

//ApiGetWxUserPhoneNumber 获取用户手机号
func (ws *WechatService) ApiGetWxUserPhoneNumber(code, token string) (wxInfo string) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=", token)
	data := bytes.NewBufferString("code=" + code)
	req, err := http.NewRequest("POST", url, data)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	wxInfo = string(body)
	log.Println(string(body))

	return
}
