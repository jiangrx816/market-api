package response

type FormatData struct {
	Value     string  `json:"value"`
	Name      string  `json:"name"`
	Checked   bool    `json:"checked"`
	Number    int     `json:"number"`
	NumberExt int     `json:"number_ext"`
	TotalFee  float64 `json:"total_fee"`
}

type MemberData struct {
	Id          int    `json:"id"`
	UserId      int64  `json:"user_id"`
	OpenId      string `json:"open_id"`
	NickName    string `json:"nick_name"`
	RealName    string `json:"real_name"`
	HeadUrl     string `json:"head_url"`
	Mobile      string `json:"mobile"`
	TagId       int    `json:"tag_id"`
	TagName     string `json:"tag_name"`
	Desc        string `json:"desc"`
	IsBest      int    `json:"is_best"`
	IsMember    int    `json:"is_member"`
	MemberLimit int    `json:"member_limit"`
	ViewCount   int    `json:"view_count"`
	Demo        string `json:"demo"`
}

type FormatTaskData struct {
	Id      int    `json:"id"`
	TagId   int    `json:"tag_id"`
	UserId  int64  `json:"user_id"`
	TagName string `json:"tag_name"`
	Desc    string `json:"desc"`
	Mobile  string `json:"mobile"`
	Date    string `json:"date"`
	Address string `json:"address"`
	Status  int    `json:"status"`
}
