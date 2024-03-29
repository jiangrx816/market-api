package model

// ZMBanner
type ZMBanner struct {
	Id     int    `json:"-"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status int    `json:"status"`
}

// TableName ZMBanner 表名
func (ZMBanner) TableName() string {
	return "zm_banner"
}

// ZMOrder
type ZMOrder struct {
	Id        int    `json:"-"`
	UserId    int64  `json:"user_id"`
	OrderId   int64  `json:"order_id"`
	Type      int    `json:"type"`
	CPrice    int    `json:"c_price"`
	OPrice    int    `json:"o_price"`
	Number    int    `json:"number"`
	NumberExt int    `json:"number_ext"`
	Status    int    `json:"status"`
	PayTime   string `json:"pay_time"`
}

// TableName ZMOrder 表名
func (ZMOrder) TableName() string {
	return "zm_order"
}

// ZMPay
type ZMPay struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CPrice    int    `json:"c_price"`
	OPrice    int    `json:"o_price"`
	Number    int    `json:"number"`
	NumberExt int    `json:"number_ext"`
	Checked   bool   `json:"checked"`
	Type      int    `json:"-"`
}

// TableName ZMPay 表名
func (ZMPay) TableName() string {
	return "zm_pay"
}

// ZMTags
type ZMTags struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Status int    `json:"-"`
}

// TableName ZMTags 表名
func (ZMTags) TableName() string {
	return "zm_tags"
}

// ZMTask
type ZMTask struct {
	Id        int    `json:"-"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	UserId    int64  `json:"user_id"`
	TagId     int    `json:"tag_id"`
	Status    int    `json:"status"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	AddTime   int64  `json:"add_time"`
}

// ZMTask ZMTask 表名
func (ZMTask) TableName() string {
	return "zm_task"
}

// ZMUser
type ZMUser struct {
	Id          int    `json:"-"`
	UserId      int64  `json:"user_id"`
	OpenId      string `json:"open_id"`
	NickName    string `json:"nick_name"`
	RealName    string `json:"real_name"`
	HeadUrl     string `json:"head_url"`
	Mobile      string `json:"mobile"`
	TagId       int    `json:"tag_id"`
	ParentId    int    `json:"parent_id"`
	IsBest      int    `json:"is_best"`
	IsMember    int    `json:"is_member"`
	MemberLimit int    `json:"member_limit"`
	Type        int    `json:"type"`
}

// ZMUser ZMUser 表名
func (ZMUser) TableName() string {
	return "zm_user"
}

// ZMUserExt
type ZMUserExt struct {
	Id        int    `json:"-"`
	UserId    int64  `json:"user_id"`
	TagId     int    `json:"tag_id"`
	Address   string `json:"address"`
	Desc      string `json:"desc"`
	Demo      string `json:"demo"`
	IsAgree   int    `json:"is_agree"`
	ViewCount int    `json:"view_count"`
}

// ZMUserExt ZMUserExt 表名
func (ZMUserExt) TableName() string {
	return "zm_user_ext"
}
