package model

// EnglishBook 英语绘本对应的书籍
type EnglishBook struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Level    uint8  `json:"level"`
	Position uint8  `json:"position"`
}

// TableName EnglishBook 表名
func (EnglishBook) TableName() string {
	return "s_english_picture"
}

// EnglishBookInfo 英语绘本对应的书籍具体的详情
type EnglishBookInfo struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	En       string `json:"en"`
	Zh       string `json:"zh"`
	Duration uint8  `json:"duration"`
	Position uint8  `json:"position"`
}

// TableName EnglishBookInfo 表名
func (EnglishBookInfo) TableName() string {
	return "s_english_picture_info"
}
