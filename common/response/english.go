package response

type ResponseEnglishBook struct {
	Id        int    `json:"id"`
	BookId    string `json:"book_id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Level     uint8  `json:"level"`
	Position  uint8  `json:"position"`
	BookCount string `json:"book_count"`
}


type OpenIdData struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}