package model

// Poetry 小学级别的古诗词对应的表
type Poetry struct {
	Id         int    `json:"-"`
	PoetryId   int    `json:"poetry_id"`
	Title      string `json:"title"`
	GradeId    uint8  `json:"grade_id"`
	Grade      string `json:"grade"`
	GradeLevel uint8  `json:"grade_level"`
	Author     string `json:"author"`
	Dynasty    string `json:"dynasty"`
	Mp3        string `json:"mp3"`
	Content    string `json:"content"`
	Info       string `json:"info"`
}

// TableName Poetry 表名
func (Poetry) TableName() string {
	return "s_school_poetry"
}

// JuniorPoetry 初高中级别的古诗词对应的表
type JuniorPoetry struct {
	Id       int    `json:"-"`
	PoetryId int    `json:"poetry_id"`
	Title    string `json:"title"`
	GradeId  uint8  `json:"grade_id"`
	Grade    string `json:"grade"`
	Author   string `json:"author"`
	Dynasty  string `json:"dynasty"`
	Content  string `json:"content"`
}

// TableName JuniorPoetry 表名
func (JuniorPoetry) TableName() string {
	return "s_junior_poetry"
}

// ChengYU 成语对应的表
type ChengYU struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Pinyin  string `json:"pinyin"`
	Explain string `json:"explain"`
	Source  string `json:"source"`
	Usage   string `json:"usage"`
	Example string `json:"example"`
	Near    string `json:"near"`
	Antonym string `json:"antonym"`
	Analyse string `json:"analyse"`
	Story   string `json:"story"`
	Level   uint8  `json:"level"`
}

// TableName ChengYU 表名
func (ChengYU) TableName() string {
	return "s_chengyu"
}