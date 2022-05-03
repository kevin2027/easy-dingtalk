package oauth2

import (
	"time"

	"github.com/kevin2027/easy-dingtalk/utils"
)

type Oauth2 interface {
	InitAccessToken(accessToken string, expireIn int64)

	GetAccessToken() (accessToken string, expireIn time.Time, err error)
	GetAgentId() (agentId int64)

	SetAgentId(agentId int64)

	GetUserToken(code string, refreshToken string) (err error)
}

func NewOuath2(opt utils.DingtalkOptions) (r Oauth2) {
	return &inner{
		appKey:    opt.AppKey,
		appSecret: opt.AppSecret,
		agentId:   opt.AgentId,
	}
}
