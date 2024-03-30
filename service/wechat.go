package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
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

type WechatService struct {
}

//ApiGetWechatData 根据code换取 openId, sessionKey, unionId
func (ws *WechatService) ApiGetWechatData(code string) (wxInfo string) {
	wechatConf := global.GVA_CONFIG.Wechat
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", wechatConf.AppId, wechatConf.Secret, code)
	if request, err := help.SendGetRequest(urlFormat); err == nil {
		wxInfo = string(request)
	}
	return
}

//ApiGetWxAccessToken 获取access_token
func (ws *WechatService) ApiGetWxAccessToken() (wxInfo string) {
	wechatConf := global.GVA_CONFIG.Wechat
	urlFormat := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wechatConf.AppId, wechatConf.Secret)
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
	wechatConf := global.GVA_CONFIG.Wechat
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/data/web/market-api/run/wx_market_cert/apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(wechatConf.MchId, wechatConf.MchCert, mchPrivateKey, wechatConf.MchIv3),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	description := "千皓优选（" + orderInfo.Name + "）"
	var cPrice int64 = int64(orderInfo.CPrice)
	svc := jsapi.JsapiApiService{Client: client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(wechatConf.AppId),
			Mchid:       core.String(wechatConf.MchId),
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

	tempOrderId := help.GetCurrentMilliseconds()
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

//ApiDealWxPayCallback 微信支付成功毁掉处理
func (ws *WechatService)ApiDealWxPayCallback(c *gin.Context) (notifyReq *notify.Request, err error) {
	ctx := c             //这个参数是context.Background()
	request := c.Request //这个值是*http.Request
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/data/web/market-api/run/wx_market_cert/apiclient_key.pem")
	if err != nil {
		log.Println("load merchant private key error")
		return nil, err
	}

	wechatConf := global.GVA_CONFIG.Wechat
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, wechatConf.MchCert, wechatConf.MchId, wechatConf.MchIv3)
	if err != nil {
		return nil, err
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(wechatConf.MchId)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(wechatConf.MchIv3, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	transaction := new(payments.Transaction)
	notifyReq, err = handler.ParseNotifyRequest(ctx, request, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		fmt.Println(err)
		//return
	}
	// 处理通知内容
	fmt.Println(notifyReq.Summary)
	fmt.Println(transaction.TransactionId)
	// 如果验签未通过，或者解密失败
	if err != nil {
		return nil, err
	}

	return notifyReq, nil
}