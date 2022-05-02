package utils

import "context"

type Attr string

const (
	AttrUserid  Attr = "userid"
	AttrUnionId Attr = "unionId"
)

type DingIdReduceFn func(ctx context.Context, attr Attr, src ...string) (dest map[string]string)

type DingtalkOptions struct {
	AppKey    string
	AppSecret string
	AgentId   int
}

type DintalkResponse struct {
	ResuestId string `json:"request_id"`
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}
