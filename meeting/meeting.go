package meeting

import (
	dingtalkconference_1_0 "github.com/alibabacloud-go/dingtalk/conference_1_0"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type Meeting interface {
	utils.DingIdReduceAble
	CreateVideoConference(unionId string, confTitle string, inviteUserIds []string, inviteCaller bool) (res *dingtalkconference_1_0.CreateVideoConferenceResponseBody, err error)
	CloseVideoConference(unionId string, conferenceId string) (err error)
	QueryConferenceInfoBatch(conferenceIdList []string) (res []*dingtalkconference_1_0.QueryConferenceInfoBatchResponseBodyInfos, err error)
}

func NewMeeting(oauth2 oauth2.Oauth2) (r Meeting) {
	return &inner{
		oauth2: oauth2,
	}
}
