package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
	"log"
	"market/common/request"
	"market/global"
	"market/model"
	help "market/utils"
	"net/http"
	"strconv"
)

var (
	mchID                      string = "1672292970"                               // 商户号
	mchCertificateSerialNumber string = "68D1E3F07BDE46784AA92001078FFF65323AE5C4" // 商户证书序列号
	mchAPIv3Key                string = "wzs920516371526000adf789cdfh9090"         // 商户APIv3密钥
)

var AppId string = "wx2afb8412b255e4fe"

type WechatService struct {
}

//ApiGetWechatData 根据code换取 openId, sessionKey, unionId
func (ws *WechatService) ApiGetWechatData(code string) (wxInfo string) {
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", global.AppId, global.AppSecret, code)
	if request, err := help.SendGetRequest(urlFormat); err == nil {
		wxInfo = string(request)
	}
	return
}

//ApiGetWxAccessToken 获取access_token
func (ws *WechatService) ApiGetWxAccessToken() (wxInfo string) {
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", global.AppId, global.AppSecret)
	if request, err := help.SendGetRequest(urlFormat); err == nil {
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
func (ws *WechatService) ApiCreateWxPay(payData request.WXPayData) (JSPayParam request.JSPayParam) {
	if payData.UserId <= 0 || payData.PayId <= 0 {
		return
	}
	orderInfo := ws.ApiCreateOrderData(payData)
	JSPayParam = ws.CreatJsApi(orderInfo)
	return JSPayParam
}

//CreatJsApi JSAPI下单
func (ws *WechatService) CreatJsApi(orderInfo model.ZMOrder) (JSPayParam request.JSPayParam) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/data/web/market-api/run/wx_market_cert/apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	//svc := certificates.CertificatesApiService{Client: client}
	//resp, result, err := svc.DownloadCertificates(ctx)
	//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)

	description := "千皓优选（" + orderInfo.Name + "）"
	var cPrice int64 = int64(orderInfo.CPrice)
	svc := jsapi.JsapiApiService{Client: client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(AppId),
			Mchid:       core.String(mchID),
			Description: core.String(description),
			OutTradeNo:  core.String(strconv.FormatInt(orderInfo.OrderId, 10)),
			Attach:      core.String("千皓优选用工好选择"),
			NotifyUrl:   core.String("https://market.58haha.com/api/wechat/pay/notice"),
			Amount: &jsapi.Amount{
				Total: core.Int64(cPrice),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(orderInfo.OpenId),
			},
		},
	)
	if err == nil {
		log.Println(resp)
		JSPayParam.PrepayID = *resp.PrepayId
		JSPayParam.Appid = *resp.Appid
		JSPayParam.TimeStamp = *resp.TimeStamp
		JSPayParam.NonceStr = *resp.NonceStr
		JSPayParam.Package = *resp.Package
		JSPayParam.SignType = *resp.SignType
		JSPayParam.PaySign = *resp.PaySign
	} else {
		log.Println(err)
	}

	return
}

//ApiCreateOrderData 生成订单信息
func (ws *WechatService) ApiCreateOrderData(payData request.WXPayData) (orderInfo model.ZMOrder) {
	//根据用户id查询用户信息
	var userInfo model.ZMUser
	global.GVA_DB.Model(&model.ZMUser{}).Debug().Where("user_id=?", payData.UserId).First(&userInfo)

	//根据payId查询相关价格
	var payItem model.ZMPay
	global.GVA_DB.Model(&model.ZMPay{}).Debug().Where("id=?", payData.PayId).First(&payItem)

	tempOrderId := help.GetCurrentUnixTimestamp()
	var order model.ZMOrder
	odb := global.GVA_DB.Model(&model.ZMOrder{}).Debug()
	order.Name = payItem.Name
	order.OpenId = userInfo.OpenId
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
