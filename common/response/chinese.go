package response

type ResponseChineseBook struct {
	Id        int    `json:"-"`
	BookId    string `json:"book_id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Level      uint8  `json:"-"`
	Position  uint8  `json:"-"`
	BookCount string `json:"book_count"`
}

type ResponseBookInfoCount struct {
	BookId    string `json:"book_id"`
	BookCount string `json:"book_count"`
}
