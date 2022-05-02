package message

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
	"golang.org/x/xerrors"
)

type Message interface {
	utils.DingIdReduceAble
	SendToConversation(userid string, cid int, msg *MessageRequest) (receiver []string, err error)

	CorpconversationaSyncsendV2(useridList []string, deptIdList []string, toAllUser bool, msg *MessageRequest) (taskId int, err error)
}

func NewMessage(oauth2 oauth2.Oauth2) (r Message) {
	return &inner{
		oauth2: oauth2,
	}
}

type inner struct {
	oauth2 oauth2.Oauth2
	utils.DingIdReduceStruct
}

func (d *inner) SendToConversation(sender string, cid int, msg *MessageRequest) (receiver []string, err error) {
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
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
	resp, err := utils.DoRquest(http.MethodPost, "/topapi/message/send_to_conversation", query, body)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var result struct {
		Receiver string `json:"receiver"`
		utils.DintalkResponse
	}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = xerrors.Errorf("%s", result.Errmsg)
		return
	}
	receiver = strings.Split(result.Receiver, "|")
	return
}
