package meeting

import (
	"context"
	"fmt"

	dingtalkconference_1_0 "github.com/alibabacloud-go/dingtalk/conference_1_0"
	"github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type inner struct {
	utils.DingIdReduceStruct
	oauth2 oauth2.Oauth2
}

func getClient() (client *dingtalkconference_1_0.Client, err error) {
	config := utils.GetOpenApiConfig()
	client, err = dingtalkconference_1_0.NewClient(config)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	return
}

func (d *inner) CreateVideoConference(unionId string, confTitle string, inviteUserIds []string, inviteCaller bool) (res *dingtalkconference_1_0.CreateVideoConferenceResponseBody, err error) {
	ids := []string{unionId}
	ids = append(ids, inviteUserIds...)

	ctx := context.Background()
	idMap := d.ReduceBatch(ctx, utils.AttrUnionId, ids...)
	if idMap == nil {
		err = fmt.Errorf("%s", "idMap is nil")
		return
	}
	if id, ok := idMap[unionId]; ok {
		unionId = id
	} else {
		err = fmt.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}
	var temp []*string
	for _, id := range inviteUserIds {
		if uid, ok := idMap[id]; ok {
			temp = append(temp, tea.String(uid))
		}
	}
	req := &dingtalkconference_1_0.CreateVideoConferenceRequest{}
	req.SetUserId(unionId)
	req.SetConfTitle(confTitle)
	req.SetInviteCaller(inviteCaller)
	if len(temp) > 0 {
		req.SetInviteUserIds(temp)
	}
	headers := &dingtalkconference_1_0.CreateVideoConferenceHeaders{}
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	headers.SetXAcsDingtalkAccessToken(accessToken)
	client, err := getClient()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	resp, err := client.CreateVideoConferenceWithOptions(req, headers, &service.RuntimeOptions{})
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	res = resp.Body
	return
}

func (d *inner) CloseVideoConference(unionId string, conferenceId string) (err error) {
	ctx := context.Background()
	unionId = d.Reduce(ctx, utils.AttrUnionId, unionId)
	if unionId == "" {
		err = fmt.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}
	req := &dingtalkconference_1_0.CloseVideoConferenceRequest{}
	req.SetUnionId(unionId)
	headers := &dingtalkconference_1_0.CloseVideoConferenceHeaders{}
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	headers.SetXAcsDingtalkAccessToken(accessToken)
	client, err := getClient()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	resp, err := client.CloseVideoConferenceWithOptions(&conferenceId, req, headers, &service.RuntimeOptions{})
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	if *resp.Body.Code != 200 {
		err = fmt.Errorf("%s", *resp.Body.Cause)
		return
	}
	return
}

func (d *inner) QueryConferenceInfoBatch(conferenceIdList []string) (res []*dingtalkconference_1_0.QueryConferenceInfoBatchResponseBodyInfos, err error) {

	headers := &dingtalkconference_1_0.QueryConferenceInfoBatchHeaders{}
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	headers.SetXAcsDingtalkAccessToken(accessToken)
	req := &dingtalkconference_1_0.QueryConferenceInfoBatchRequest{}
	client, err := getClient()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	resp, err := client.QueryConferenceInfoBatchWithOptions(req, headers, &service.RuntimeOptions{})
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	res = resp.Body.Infos
	return
}
