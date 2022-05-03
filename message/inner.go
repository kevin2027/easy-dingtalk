package message

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type inner struct {
	oauth2 oauth2.Oauth2
	utils.DingIdReduceStruct
}

func (d *inner) SendToConversation(sender string, cid int, msg *MessageRequest) (receiver []string, err error) {
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	ctx := context.Background()
	sender = d.Reduce(ctx, utils.AttrUserid, sender)
	body := make(map[string]interface{})
	body["sender"] = sender
	body["cid"] = cid
	body["msg"] = msg.Clone()

	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRequest(http.MethodPost, "/topapi/message/send_to_conversation", query, body)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	var result struct {
		Receiver string `json:"receiver"`
		utils.DintalkResponse
	}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	receiver = strings.Split(result.Receiver, "|")
	return
}
