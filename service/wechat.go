package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"market/common/request"
	"market/global"
	"market/model"
	"market/utils"
	"net/http"
	"strconv"
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
func (ws *WechatService) ApiGetWxUserPhoneNumber(photoData request.MakePhotoData) (wxInfo string) {
	if photoData.Token == "" || photoData.Code == "" {
		return
	}
	// 构造请求的数据
	requestBody, err := json.Marshal(map[string]string{
		"code": photoData.Code,
	})
	if err != nil {
		fmt.Println("Error encoding request body:", err)
		return
	}

	// 发起 POST 请求
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", photoData.Token), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	wxInfo = string(body)
	return
}

//ApiGetWxPay 创建订单
func (ws *WechatService) ApiCreateWxPay(payData request.WXPayData) {
	if payData.UserId <= 0 || payData.PayId <= 0 {
		return
	}
	orderInfo := ws.ApiCreateOrderData(payData)
	ws.CreatJsApi(orderInfo)
}

//CreatJsApi JSAPI下单
func (ws *WechatService) CreatJsApi(orderInfo model.ZMOrder) {
	url := "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"
	var param request.PayDataParams
	param.Mchid = "1672292970"
	param.Appid = "wx2afb8412b255e4fe"
	param.OutTradeNo = strconv.FormatInt(orderInfo.OrderId, 10)
	param.Description = orderInfo.Name
	param.NotifyURL = "https://market.58haha.com/api/wechat/pay/notice"
	param.Amount.Total = orderInfo.CPrice
	param.Amount.Currency = "CNY"
	param.Payer.Openid = orderInfo.OpenId

	marshal, err2 := json.Marshal(param)
	if err2 != nil {
		return
	}
	jsonStr := []byte(marshal)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Authorization", "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 解析响应
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//ApiCreateOrderData 生成订单信息
func (ws *WechatService) ApiCreateOrderData(payData request.WXPayData) (orderInfo model.ZMOrder) {
	//根据用户id查询用户信息
	var userInfo model.ZMUser
	global.GVA_DB.Model(&model.ZMUser{}).Debug().Where("user_id=?", payData.UserId).First(&userInfo)

	//根据payId查询相关价格
	var payItem model.ZMPay
	global.GVA_DB.Model(&model.ZMPay{}).Debug().Where("id=?", payData.PayId).First(&payItem)

	tempOrderId := utils.GetCurrentUnixTimestamp()
	var order model.ZMOrder
	odb := global.GVA_DB.Model(&model.ZMOrder{}).Debug()
	order.Name = payItem.Name
	order.UserId = userInfo.UserId
	order.OrderId = tempOrderId
	order.Type = 1 //1普通会员,2优选工匠,3积分兑换
	order.CPrice = payItem.CPrice
	order.OPrice = payItem.OPrice
	order.Number = payItem.Number
	order.NumberExt = payItem.NumberExt
	order.Status = 0 //支付状态,1支付完成,0待支付
	affected := odb.Create(&order).RowsAffected
	if affected > 0 {
		orderInfo = order
	}
	return
}
