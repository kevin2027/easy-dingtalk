package oauth2

import (
	"time"

	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type Oauth2 interface {
	InitAccessToken(accessToken string, expireIn int64)

	GetAccessToken() (accessToken string, expireIn time.Time, err error)
	GetAgentId() (agentId int64)

	SetAgentId(agentId int64)

	GetUserToken(code string, refreshToken string) (res *dingtalkoauth2_1_0.GetUserTokenResponseBody, err error)
}

func NewOuath2(opt utils.DingtalkOptions) (r Oauth2) {
	return &inner{
		appKey:    opt.AppKey,
		appSecret: opt.AppSecret,
		agentId:   opt.AgentId,
	}
}
