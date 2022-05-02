package message

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/utils"
	"golang.org/x/xerrors"
)

func (d *inner) CorpconversationaSyncsendV2(useridList []string, deptIdList []string, toAllUser bool, msg *MessageRequest) (taskId int, err error) {
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if len(useridList) > 0 {
		ctx := context.Background()
		userIdMap := d.ReduceBatch(ctx, utils.AttrUserid, useridList...)
		temp := make([]string, 0, len(useridList))
		for _, id := range useridList {
			if userid, ok := userIdMap[id]; ok {
				temp = append(temp, userid)
			}
		}
		useridList = temp
	}
	if len(deptIdList) > 0 {
		ctx := context.Background()
		deptIdMap := d.ReduceBatch(ctx, utils.AttrUserid, deptIdList...)
		temp := make([]string, 0, len(useridList))
		for _, id := range deptIdList {
			if userid, ok := deptIdMap[id]; ok {
				temp = append(temp, userid)
			}
		}
		deptIdList = temp
	}

	body := make(map[string]interface{})
	body["agent_id"] = d.oauth2.GetAgentId()
	body["to_all_user"] = toAllUser
	if len(useridList) > 0 {
		body["userid_list"] = strings.Join(useridList, ",")
	}
	if len(deptIdList) > 0 {
		body["dept_id_list"] = strings.Join(deptIdList, ",")
	}
	body["msg"] = msg.Clone()

	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRquest(http.MethodPost, "/topapi/message/corpconversation/asyncsend_v2", query, body)
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
		TaskId int `json:"task_id"`
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
	taskId = result.TaskId
	return
}
