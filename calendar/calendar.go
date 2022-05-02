package calendar

import (
	"context"

	dingtalkcalendar_1_0 "github.com/alibabacloud-go/dingtalk/calendar_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/contact"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
	"golang.org/x/xerrors"
)

type Calendar interface {
	CreateEvent(userId string, req *dingtalkcalendar_1_0.CreateEventRequest) (event *dingtalkcalendar_1_0.CreateEventResponseBody, err error)

	PatchEvent(userId string, eventId string, req *dingtalkcalendar_1_0.PatchEventRequest) (event *dingtalkcalendar_1_0.PatchEventResponseBody, err error)

	DeleteEvent(userId string, eventId string) (err error)

	AddAttendee(userId string, eventId string, req *dingtalkcalendar_1_0.AddAttendeeRequest) (err error)

	RemoveAttendee(userId string, eventId string, req *dingtalkcalendar_1_0.RemoveAttendeeRequest) (err error)

	SetDingDiReduceFn(fn utils.DingIdReduceFn)
}

func NewCalendar(oauth2 oauth2.Oauth2,
	contact contact.Contact,
) (r Calendar) {
	return &inner{
		oauth2:  oauth2,
		contact: contact,
	}
}

type inner struct {
	oauth2         oauth2.Oauth2
	contact        contact.Contact
	dingIdReduceFn utils.DingIdReduceFn
}

func getClient() (client *dingtalkcalendar_1_0.Client, err error) {
	config := utils.GetOpenApiConfig()
	client, err = dingtalkcalendar_1_0.NewClient(config)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (d *inner) SetDingDiReduceFn(fn utils.DingIdReduceFn) {
	d.dingIdReduceFn = fn
}

func (d *inner) CreateEvent(userId string, req *dingtalkcalendar_1_0.CreateEventRequest) (event *dingtalkcalendar_1_0.CreateEventResponseBody, err error) {

	var attendees []string
	for _, a := range req.Attendees {
		attendees = append(attendees, *a.Id)
	}
	attendees = append(attendees, userId)
	ctx := context.Background()
	attendeeMap := utils.DingIdReduceBatch(d.dingIdReduceFn, ctx, attendees...)
	if attendeeMap == nil {
		err = xerrors.Errorf("%s", "attendeeMap is nil")
		return
	}
	if id, ok := attendeeMap[userId]; ok {
		userId = id
	} else {
		err = xerrors.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}

	attendeeList := make([]*dingtalkcalendar_1_0.CreateEventRequestAttendees, 0, len(attendees))
	for _, attendee := range req.Attendees {
		if id, ok := attendeeMap[*attendee.Id]; ok {
			attendee.Id = tea.String(id)
			attendeeList = append(attendeeList, attendee)
		}
	}
	if len(attendeeList) > 0 {
		req.SetAttendees(attendeeList)
	} else {
		req.SetAttendees(nil)
	}

	client, err := getClient()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	createEventHeaders := &dingtalkcalendar_1_0.CreateEventHeaders{}
	createEventHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	res, err := client.CreateEventWithOptions(tea.String(userId), tea.String("primary"), req, createEventHeaders, &util.RuntimeOptions{})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	event = res.Body
	return

}

func (d *inner) PatchEvent(userId string, eventId string, req *dingtalkcalendar_1_0.PatchEventRequest) (event *dingtalkcalendar_1_0.PatchEventResponseBody, err error) {
	var attendees []string
	for _, a := range req.Attendees {
		attendees = append(attendees, *a.Id)
	}
	attendees = append(attendees, userId)
	ctx := context.Background()
	attendeeMap := utils.DingIdReduceBatch(d.dingIdReduceFn, ctx, attendees...)
	if attendeeMap == nil {
		err = xerrors.Errorf("%s", "attendeeMap is nil")
		return
	}
	if id, ok := attendeeMap[userId]; ok {
		userId = id
	} else {
		err = xerrors.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}

	attendeeList := make([]*dingtalkcalendar_1_0.PatchEventRequestAttendees, 0, len(attendees))
	for _, attendee := range req.Attendees {
		if id, ok := attendeeMap[*attendee.Id]; ok {
			attendee.Id = tea.String(id)
			attendeeList = append(attendeeList, attendee)
		}
	}
	if len(attendeeList) > 0 {
		req.SetAttendees(attendeeList)
	} else {
		req.SetAttendees(nil)
	}
	client, err := getClient()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	headers := &dingtalkcalendar_1_0.PatchEventHeaders{}
	headers.XAcsDingtalkAccessToken = tea.String(accessToken)
	userInfo, err := d.contact.GetUserInfo(userId)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	userId = userInfo.Unionid
	res, err := client.PatchEventWithOptions(tea.String(userId), tea.String("primary"), tea.String(eventId), req, headers, &util.RuntimeOptions{})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	event = res.Body
	return
}

func (d *inner) DeleteEvent(userId string, eventId string) (err error) {
	ctx := context.Background()
	userId = utils.DingIdReduce(d.dingIdReduceFn, ctx, userId)
	if userId == "" {
		err = utils.ErrUserIdIsEmpty
		return
	}
	client, err := getClient()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	headers := &dingtalkcalendar_1_0.DeleteEventHeaders{}
	headers.XAcsDingtalkAccessToken = tea.String(accessToken)
	_, err = client.DeleteEventWithOptions(tea.String(userId), tea.String("primary"), tea.String(eventId), headers, &util.RuntimeOptions{})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (d *inner) AddAttendee(userId string, eventId string, req *dingtalkcalendar_1_0.AddAttendeeRequest) (err error) {
	var attendees []string
	for _, a := range req.AttendeesToAdd {
		attendees = append(attendees, *a.Id)
	}
	attendees = append(attendees, userId)
	ctx := context.Background()
	attendeeMap := utils.DingIdReduceBatch(d.dingIdReduceFn, ctx, attendees...)
	if attendeeMap == nil {
		err = xerrors.Errorf("%s", "attendeeMap is nil")
		return
	}
	if id, ok := attendeeMap[userId]; ok {
		userId = id
	} else {
		err = xerrors.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}

	attendeeList := make([]*dingtalkcalendar_1_0.AddAttendeeRequestAttendeesToAdd, 0, len(attendees))
	for _, attendee := range req.AttendeesToAdd {
		if id, ok := attendeeMap[*attendee.Id]; ok {
			attendee.Id = tea.String(id)
			attendeeList = append(attendeeList, attendee)
		}
	}
	if len(attendeeList) > 0 {
		req.SetAttendeesToAdd(attendeeList)
	} else {
		req.SetAttendeesToAdd(nil)
	}
	client, err := getClient()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	headers := &dingtalkcalendar_1_0.AddAttendeeHeaders{}
	headers.XAcsDingtalkAccessToken = tea.String(accessToken)
	_, err = client.AddAttendeeWithOptions(tea.String(userId), tea.String("primary"), tea.String(eventId), req, headers, &util.RuntimeOptions{})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
func (d *inner) RemoveAttendee(userId string, eventId string, req *dingtalkcalendar_1_0.RemoveAttendeeRequest) (err error) {
	var attendees []string
	for _, a := range req.AttendeesToRemove {
		attendees = append(attendees, *a.Id)
	}
	attendees = append(attendees, userId)
	ctx := context.Background()
	attendeeMap := utils.DingIdReduceBatch(d.dingIdReduceFn, ctx, attendees...)
	if attendeeMap == nil {
		err = xerrors.Errorf("%s", "attendeeMap is nil")
		return
	}
	if id, ok := attendeeMap[userId]; ok {
		userId = id
	} else {
		err = xerrors.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}

	attendeeList := make([]*dingtalkcalendar_1_0.RemoveAttendeeRequestAttendeesToRemove, 0, len(attendees))
	for _, attendee := range req.AttendeesToRemove {
		if id, ok := attendeeMap[*attendee.Id]; ok {
			attendee.Id = tea.String(id)
			attendeeList = append(attendeeList, attendee)
		}
	}
	if len(attendeeList) > 0 {
		req.SetAttendeesToRemove(attendeeList)
	} else {
		req.SetAttendeesToRemove(nil)
	}
	client, err := getClient()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	headers := &dingtalkcalendar_1_0.RemoveAttendeeHeaders{}
	headers.XAcsDingtalkAccessToken = tea.String(accessToken)
	_, err = client.RemoveAttendeeWithOptions(tea.String(userId), tea.String("primary"), tea.String(eventId), req, headers, &util.RuntimeOptions{})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
