package model

// Category 结构体
type Category struct {
	Id   int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	Name string `json:"name" form:"name" gorm:"column:name;comment:栏目名称;size:255;"`
}

// TableName Category 表名
func (Category) TableName() string {
	return "s_category"
}

// BaiKe 结构体
type BaiKe struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Question   string `json:"question"`
	OptionA    string `json:"option_a"`
	OptionB    string `json:"option_b"`
	OptionC    string `json:"option_c"`
	OptionD    string `json:"option_d"`
	Answer     string `json:"answer"`
	Analytic   string `json:"analytic"`
	AddTime    string `json:"add_time"`
}

// TableName BaiKe 表名
func (BaiKe) TableName() string {
	return "s_baike"
}
