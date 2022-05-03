package calendar_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type inner struct {
	oauth2 oauth2.Oauth2
	utils.DingIdReduceStruct
}

func (d *inner) CreateEvent(event *CreateEventRequestEvent) (res *CreateEventResponseResult, err error) {
	var attendees []string
	for _, a := range event.Attendees {
		attendees = append(attendees, *a.Userid)
	}
	attendees = append(attendees, *event.Organizer.Userid)
	ctx := context.Background()
	attendeesMap := d.ReduceBatch(ctx, utils.AttrUserid, attendees...)

	userId := attendeesMap[*event.Organizer.Userid]
	if userId == "" {
		err = fmt.Errorf("%w", utils.ErrUserIdIsEmpty)
		return
	}
	event.Organizer.Userid = tea.String(userId)

	attendeeList := make([]*Attendee, 0, len(attendees))
	for _, attendee := range event.Attendees {
		if id, ok := attendeesMap[*attendee.Userid]; ok {
			if id != "" {
				attendeeList = append(attendeeList, &Attendee{
					Userid:         tea.String(id),
					AttendeeStatus: attendee.AttendeeStatus,
				})
			}
		}
	}
	event.Attendees = attendeeList
	event.CalendarId = "primary"
	event.NotificationType = "NONE"
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	body := make(map[string]interface{})
	body["agentid"] = d.oauth2.GetAgentId()
	body["event"] = *event
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}

	resp, err := utils.DoRequest(http.MethodPost, "/topapi/calendar/v2/event/create", query, body)
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
	var result CreateEventResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = fmt.Errorf("%s", "success is false")
		return
	}
	res = result.Result
	return
}

func (d *inner) UpdateEvent(event *UpdateEventRequestEvent) (err error) {
	var attendees []string
	for _, a := range event.Attendees {
		attendees = append(attendees, *a.Userid)
	}
	attendees = append(attendees, *event.Organizer.Userid)
	ctx := context.Background()
	attendeesMap := d.ReduceBatch(ctx, utils.AttrUserid, attendees...)

	userId := attendeesMap[*event.Organizer.Userid]
	if userId == "" {
		err = utils.ErrUserIdIsEmpty
		return
	}
	event.Organizer.Userid = tea.String(userId)

	attendeeList := make([]*Attendee, 0, len(attendees))
	for _, attendee := range event.Attendees {
		if id, ok := attendeesMap[*attendee.Userid]; ok {
			if id != "" {
				attendeeList = append(attendeeList, &Attendee{
					Userid:         tea.String(id),
					AttendeeStatus: attendee.AttendeeStatus,
				})
			}
		}
	}
	event.Attendees = attendeeList
	event.CalendarId = "primary"
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}

	body := make(map[string]interface{})
	body["agentid"] = d.oauth2.GetAgentId()
	body["event"] = *event
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRequest(http.MethodPost, "/topapi/calendar/v2/event/update", query, body)
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
	var result DintalkSuccessResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = fmt.Errorf("%s", "success is false")
		return
	}
	return
}

func (d *inner) CancelEvent(eventId string) (err error) {
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	body := make(map[string]interface{})
	body["event_id"] = eventId
	body["calendar_id"] = "primary"
	body["agentid"] = d.oauth2.GetAgentId()
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRequest(http.MethodPost, "/topapi/calendar/v2/event/cancel", query, body)
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
	var result DintalkSuccessResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = fmt.Errorf("%s", "success is false")
		return
	}
	return
}

func (d *inner) AttendeeUpdate(eventId string, eventAtttendees []*Attendee) (err error) {
	var attendees []string
	for _, a := range eventAtttendees {
		attendees = append(attendees, *a.Userid)
	}
	ctx := context.Background()
	attendeesMap := d.ReduceBatch(ctx, utils.AttrUserid, attendees...)

	attendeeList := make([]*Attendee, 0, len(attendees))
	for _, attendee := range eventAtttendees {
		if id, ok := attendeesMap[*attendee.Userid]; ok {
			if id != "" {
				attendeeList = append(attendeeList, &Attendee{
					Userid:         tea.String(id),
					AttendeeStatus: attendee.AttendeeStatus,
				})
			}
		}
	}
	accessToken, _, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}

	body := make(map[string]interface{})
	body["agentid"] = d.oauth2.GetAgentId()
	body["calendar_id"] = "primary"
	body["event_id"] = eventId
	body["attendees"] = attendeeList

	resp, err := utils.DoRequest(http.MethodPost, "/topapi/calendar/v2/attendee/update", query, body)
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
	var result DintalkSuccessResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = fmt.Errorf("%s", "success is false")
		return
	}
	return
}
