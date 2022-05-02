package utils

import "context"

type GetOauth2Fn func() (accessToken string, err error)

type DingIdReduceFn func(ctx context.Context, src ...string) (dest map[string]string)

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
