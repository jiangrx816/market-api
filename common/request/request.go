package request

type MakeTaskData struct {
	TaskDesc  string `json:"task_desc"`
	TagId  int `json:"tag_id"`
	UserId  int `json:"user_id"`
	Address  string `json:"address"`
	Title  string `json:"title"`
}
