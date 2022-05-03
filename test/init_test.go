package dingtalk_test

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kevin2027/easy-dingtalk/dingtalk"
	"github.com/kevin2027/easy-dingtalk/utils"
	"github.com/kevin2027/merrors"
)

//go:embed .cfg/config.json
var configData []byte
var config = struct {
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	AgentId     int64  `json:"agent_id"`
	AccessToken string `json:"access_token"`
	ExpireIn    int64  `json:"expire_in"`
	Users       map[string]struct {
		Name   string `json:"name"`
		Userid string `json:"userid"`
	} `json:"users"`
}{}

var client dingtalk.Dingtalk

func deferErr(err *error) {
	if *err != nil {
		fmt.Printf("%+v\n", merrors.String(*err))
	}
}

func init() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("%+v\n", merrors.String(err))
			os.Exit(1)
		}
	}()
	fmt.Printf("init ....\n")
	err = json.Unmarshal(configData, &config)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	client, _, err = dingtalk.NewDingtalk(utils.DingtalkOptions{
		AppKey:    config.AppKey,
		AppSecret: config.AppSecret,
		AgentId:   config.AgentId,
	})
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	client.SetDingDiReduceFn(func(ctx context.Context, attr utils.Attr, src ...string) (dest map[string]string) {
		dest = make(map[string]string)
		if attr == utils.AttDeptId {
			return
		}
		for _, s := range src {
			if user, ok := config.Users[s]; ok {
				switch attr {
				case utils.AttrUserid:
					dest[s] = user.Userid
				}
			}
		}

		return
	})
	now := time.Now().Unix()
	if config.AccessToken != "" && config.ExpireIn > now+300 {
		fmt.Printf("旧 accessToken: %s\n", config.AccessToken)
		client.Oauth2().InitAccessToken(config.AccessToken, config.ExpireIn)
	} else {
		var accessToken string
		var expireIn time.Time
		accessToken, expireIn, err = client.Oauth2().GetAccessToken()
		if err != nil {
			err = fmt.Errorf("%w", err)
			return
		}
		config.AccessToken = accessToken
		fmt.Printf("新 accessToken: %s\n", config.AccessToken)
		config.ExpireIn = expireIn.Unix()
		var data []byte
		data, err = json.MarshalIndent(&config, "", "  ")
		if err != nil {
			err = fmt.Errorf("%w", err)
			return
		}
		err = ioutil.WriteFile(".cfg/config.json", data, 0644)
		if err != nil {
			err = fmt.Errorf("%w", err)
			return
		}
	}
	fmt.Printf("init end\n")
}
