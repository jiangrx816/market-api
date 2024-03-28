package request

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
