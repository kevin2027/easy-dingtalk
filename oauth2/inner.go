package oauth2

import (
	"fmt"
	"sync"
	"time"

	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type inner struct {
	accessTokenMutex struct {
		sync.RWMutex
		accessToken string
		expireIn    time.Time
	}
	jsapiTicketMutex struct {
		sync.RWMutex
		jsapiTicket string
		expireIn    time.Time
	}
	appKey    string
	appSecret string
	agentId   int64
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
	d.accessTokenMutex.Lock()
	defer d.accessTokenMutex.Unlock()
	d.accessTokenMutex.accessToken = accessToken
	d.accessTokenMutex.expireIn = time.Unix(expireIn, 0)

}

func (d *inner) GetAgentId() (agentId int64) {
	return d.agentId
}

func (d *inner) SetAgentId(agentId int64) {
	d.agentId = agentId
}

func (d *inner) GetAccessToken() (accessToken string, expireIn time.Time, err error) {
	d.accessTokenMutex.RLock()
	if d.accessTokenMutex.accessToken != "" && d.accessTokenMutex.expireIn.After(time.Now()) {
		accessToken = d.accessTokenMutex.accessToken
		expireIn = d.accessTokenMutex.expireIn
		d.accessTokenMutex.RUnlock()
		return
	}
	d.accessTokenMutex.RUnlock()
	d.accessTokenMutex.Lock()
	defer d.accessTokenMutex.Unlock()
	if d.accessTokenMutex.accessToken != "" && d.accessTokenMutex.expireIn.After(time.Now()) {
		accessToken = d.accessTokenMutex.accessToken
		expireIn = d.accessTokenMutex.expireIn
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
	d.accessTokenMutex.accessToken = *res.Body.AccessToken
	d.accessTokenMutex.expireIn = time.Now().Add(time.Duration(*res.Body.ExpireIn)*time.Second - time.Minute)
	accessToken = d.accessTokenMutex.accessToken
	expireIn = d.accessTokenMutex.expireIn
	return
}

func (d *inner) GetUserToken(code string, refreshToken string) (res *dingtalkoauth2_1_0.GetUserTokenResponseBody, err error) {
	client, err := getClient()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	req := &dingtalkoauth2_1_0.GetUserTokenRequest{
		ClientId:     tea.String(d.appKey),
		ClientSecret: tea.String(d.appSecret),
		Code:         tea.String(code),
		RefreshToken: nil,
		GrantType:    tea.String("authorization_code"),
	}
	if len(refreshToken) > 0 {
		req.RefreshToken = tea.String(refreshToken)
		req.GrantType = tea.String("refresh_token")
	}
	resp, err := client.GetUserToken(req)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	res = resp.Body
	return
}

func (d *inner) CreateJsapiTicket() (jsapiTicket string, expireIn time.Time, err error) {
	d.jsapiTicketMutex.RLock()
	if d.jsapiTicketMutex.jsapiTicket != "" && d.jsapiTicketMutex.expireIn.After(time.Now()) {
		jsapiTicket = d.jsapiTicketMutex.jsapiTicket
		expireIn = d.jsapiTicketMutex.expireIn
		d.jsapiTicketMutex.RUnlock()
		return
	}
	d.jsapiTicketMutex.RUnlock()
	d.jsapiTicketMutex.Lock()
	defer d.jsapiTicketMutex.Unlock()
	if d.jsapiTicketMutex.jsapiTicket != "" && d.jsapiTicketMutex.expireIn.After(time.Now()) {
		jsapiTicket = d.jsapiTicketMutex.jsapiTicket
		expireIn = d.jsapiTicketMutex.expireIn
		return
	}
	client, err := getClient()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	accessToken, _, err := d.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	headers := &dingtalkoauth2_1_0.CreateJsapiTicketHeaders{
		XAcsDingtalkAccessToken: tea.String(accessToken),
	}
	resp, err := client.CreateJsapiTicketWithOptions(headers, &service.RuntimeOptions{})
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	d.jsapiTicketMutex.jsapiTicket = *resp.Body.JsapiTicket
	d.jsapiTicketMutex.expireIn = time.Now().Add(time.Duration(*resp.Body.ExpireIn)*time.Second - time.Minute)
	jsapiTicket = d.jsapiTicketMutex.jsapiTicket
	expireIn = d.jsapiTicketMutex.expireIn
	return
}
