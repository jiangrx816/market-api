package response

type ResponseSchoolPoetry struct {
	Id         int    `json:"id"`
	PoetryId   int    `json:"poetry_id"`
	Title      string `json:"title"`
	GradeId    uint8  `json:"grade_id" `
	Grade      string `json:"grade" `
	GradeLevel uint8  `json:"grade_level" `
	Author     string `json:"author" `
	Dynasty    string `json:"dynasty"`
}

type ResponseSchoolPoetryData struct {
	Id             int          `json:"id"`
	PoetryId       int          `json:"poetry_id"`
	Title          string       `json:"title"`
	GradeId        uint8        `json:"grade_id"`
	Grade          string       `json:"grade"`
	GradeLevel     uint8        `json:"grade_level"`
	Author         string       `json:"author"`
	Dynasty        string       `json:"dynasty" `
	Mp3            string       `json:"mp3" `
	Content        string       `json:"content"`
	Info           string       `json:"info"`
	ListContent    []string     `json:"list_content"`
	ListInfo       []string     `json:"list_info"`
	PoetryListInfo []PoetryInfo `json:"poetry_list_info"`
}

type PoetryInfo struct {
	ZH string `json:"zh"`
	PY string `json:"py"`
}

type CYdATA struct {
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
	StoryList []string `json:"story_list"`
}