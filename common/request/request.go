package request

type UpdateTaskStatus struct {
	TaskId int `json:"task_id"`
	Status int `json:"status"`
}

type MakeTaskData struct {
	TaskDesc string `json:"task_desc"`
	TagId    int    `json:"tag_id"`
	UserId   int64  `json:"user_id"`
	Address  string `json:"address"`
	Title    string `json:"title"`
}

type MakeUserData struct {
	Mobile   string `json:"mobile"`
	OpenId   string `json:"open_id"`
	Type     int    `json:"type"`
	NickName string `json:"nick_name"`
	HeadImg  string `json:"head_img"`
}

type MemberUpdateData struct {
	UserId   int    `json:"user_id"`
	NickName string `json:"nick_name"`
	HeadUrl  string `json:"head_url"`
	Mobile   string `json:"mobile"`
}

type MakePhotoData struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}

type WXPayData struct {
	UserId int `json:"user_id"`
	PayId  int `json:"pay_id"`
}

type PayDataParams struct {
	Mchid       string `json:"mchid"`
	OutTradeNo  string `json:"out_trade_no"`
	Appid       string `json:"appid"`
	Description string `json:"description"`
	NotifyURL   string `json:"notify_url"`
	Amount      struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Payer struct {
		Openid string `json:"openid"`
	} `json:"payer"`
}

type JSPayParam struct {
	PrepayID  string `json:"PrepayId"`
	Appid     string `json:"Appid"`
	TimeStamp string `json:"TimeStamp"`
	NonceStr  string `json:"NonceStr"`
	Package   string `json:"Package"`
	SignType  string `json:"SignType"`
	PaySign   string `json:"PaySign"`
}
