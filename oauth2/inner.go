package oauth2

import (
	"fmt"
	"sync"
	"time"

	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type inner struct {
	sync.RWMutex
	accessToken string
	expireIn    time.Time
	appKey      string
	appSecret   string
	agentId     int64
}

func getClient() (client *dingtalkoauth2_1_0.Client, err error) {
	config := utils.GetOpenApiConfig()
	client, err = dingtalkoauth2_1_0.NewClient(config)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	return
}
func (d *inner) InitAccessToken(accessToken string, expireIn int64) {
	d.Lock()
	defer d.Unlock()
	d.accessToken = accessToken
	d.expireIn = time.Unix(expireIn, 0)

}

func (d *inner) GetAgentId() (agentId int64) {
	return d.agentId
}

func (d *inner) SetAgentId(agentId int64) {
	d.agentId = agentId
}

func (d *inner) GetAccessToken() (accessToken string, expireIn time.Time, err error) {
	d.RLock()
	if d.accessToken != "" && d.expireIn.After(time.Now()) {
		accessToken = d.accessToken
		expireIn = d.expireIn
		d.RUnlock()
		return
	}
	d.RUnlock()
	d.Lock()
	defer d.Unlock()
	if d.accessToken != "" && d.expireIn.After(time.Now()) {
		accessToken = d.accessToken
		return
	}
	client, err := getClient()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	res, err := client.GetAccessToken(&dingtalkoauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(d.appKey),
		AppSecret: tea.String(d.appSecret),
	})
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	d.accessToken = *res.Body.AccessToken
	d.expireIn = time.Now().Add(time.Duration(*res.Body.ExpireIn)*time.Second - time.Minute)
	accessToken = d.accessToken
	expireIn = d.expireIn
	return
}

func (d *inner) GetUserToken(code string, refreshToken string) (err error) {
	// client, err := getClient()
	// if err != nil {
	// 	err = fmt.Errorf("%w", err)
	// 	return
	// }
	return
}
