package utils

type DingtalkOptions struct {
	AppKey    string
	AppSecret string
	AgentId   int64
}

type DintalkResponse struct {
	ResuestId string `json:"request_id"`
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}
