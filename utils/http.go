package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HttpParams struct {
	Param string `json:"param"`
	Item  string `json:"item"`
}

//SendHttpPost 发起http的post方式请求
func SendHttpPost(httpUrl string, HttpParamList []HttpParams) (body []byte) {
	// 构造要发送的参数
	data := url.Values{}

	for idx, _ := range HttpParamList {
		data.Set(HttpParamList[idx].Param, HttpParamList[idx].Item)
	}

	// 创建HTTP请求的客户端
	client := &http.Client{}

	// 构造POST请求
	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 处理响应
	// 在这里可以读取resp.Body来获取服务器的响应数据
	body, _ = ioutil.ReadAll(resp.Body)

	//返回json，去指定的地方处理
	//jsonString = string(body)
	return
}
