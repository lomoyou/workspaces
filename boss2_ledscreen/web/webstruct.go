package web

type postinfo struct {
	User string `json:"user"`
	Sign string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Data string `json:"data"`
	Version string `json:"version"`
	Method string `json:"method"`
	Requesttoken string `json:"requesttoken"`
}

type carinfo struct {
	Packcode string `json:"packcode"`
	Total_num int `json:"total_num"`
	Empty_num int `json:"empty_num"`
	Used_num int `json:"used_num"`
}

type resultinfo struct {
	Success bool `json:"success"`
	Code int `json:"code"`
	Message string `json:"message"`
	Data string `json:"data"`
	Requesttoken string `json:"requesttoken"`
	Sign string `json:"sign"`
}