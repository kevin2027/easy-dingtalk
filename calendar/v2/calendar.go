package calendar_v2

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kevin2027/easy-dingtalk/oauth2"
	"github.com/Kevin2027/easy-dingtalk/utils"
	"github.com/alibabacloud-go/tea/tea"
	"golang.org/x/xerrors"
)

type Calendar interface {
	CreateEvent(event *CreateEventRequestEvent) (res *CreateEventResponseResult, err error)
	UpdateEvent(event *UpdateEventRequestEvent) (err error)
	CancelEvent(eventId string) (err error)

	SetDingDiReduceFn(fn utils.DingIdReduceFn)
}

func NewCalendar(oauth2 oauth2.Oauth2) (r Calendar) {
	return &inner{
		oauth2: oauth2,
	}
}

type inner struct {
	oauth2         oauth2.Oauth2
	dingIdReduceFn utils.DingIdReduceFn
}

type CreateEventRequestEvent struct {
	Organizer        Attendee    `json:"organizer"`
	CalendarId       string      `json:"calendar_id"`
	Attendees        []*Attendee `json:"attendees,omitempty"`
	Summary          string      `json:"summary"`
	Description      *string     `json:"description,omitempty"`
	Start            DataTime    `json:"start"`
	End              DataTime    `json:"end"`
	Reminder         *Reminder   `json:"reminder,omitempty"`
	Location         *Location   `json:"location,omitempty"`
	NotificationType string      `json:"notification_type,omitempty"`
}

type CreateEventResponse struct {
	DintalkSuccessResponse
	Result *CreateEventResponseResult `json:"result"`
}

type CreateEventResponseResult struct {
	Attendees   []*Attendee `json:"attendees"`
	CalendarId  string      `json:"calendar_id"`
	Description string      `json:"description"`
	End         *DataTime   `json:"end"`
	EventId     string      `json:"event_id"`
	Organizer   *Attendee   `json:"organizer"`
	Start       *DataTime   `json:"start"`
	Summary     string      `json:"summary"`
	Reminder    *Reminder   `json:"reminder"`
	Location    *Location   `json:"location"`
}

func (d *inner) SetDingDiReduceFn(fn utils.DingIdReduceFn) {
	d.dingIdReduceFn = fn
}

func (d *inner) CreateEvent(event *CreateEventRequestEvent) (res *CreateEventResponseResult, err error) {
	var attendees []string
	for _, a := range event.Attendees {
		attendees = append(attendees, *a.Userid)
	}
	attendees = append(attendees, *event.Organizer.Userid)
	ctx := context.Background()
	attendeesMap := utils.DingIdReduceBatch(d.dingIdReduceFn, ctx, attendees...)

	userId := attendeesMap[*event.Organizer.Userid]
	if userId == "" {
		err = xerrors.Errorf("%w", utils.ErrUserIdIsEmpty)
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
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	body := make(map[string]interface{})
	body["agentid"] = d.oauth2.GetAgentId()
	body["event"] = *event
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}

	resp, err := utils.DoRquest(http.MethodPost, "/topapi/calendar/v2/event/create", query, body)
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
	var result CreateEventResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = xerrors.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = xerrors.Errorf("%s", "success is false")
		return
	}
	res = result.Result
	return
}

type UpdateEventRequestEvent struct {
	Attendees   []*Attendee `json:"attendees,omitempty"`
	CalendarId  string      `json:"calendar_id"`
	Description string      `json:"description"`
	End         DataTime    `json:"end"`
	Start       DataTime    `json:"start"`
	Summary     string      `json:"summary"`
	EventId     string      `json:"event_id"`
	Reminder    *Reminder   `json:"reminder,omitempty"`
	Location    *Location   `json:"location,omitempty"`
	Organizer   Attendee    `json:"organizer"`
}

func (d *inner) UpdateEvent(event *UpdateEventRequestEvent) (err error) {
	var attendees []string
	for _, a := range event.Attendees {
		attendees = append(attendees, *a.Userid)
	}
	attendees = append(attendees, *event.Organizer.Userid)
	ctx := context.Background()
	attendeesMap := utils.DingIdReduceBatch(d.dingIdReduceFn, ctx, attendees...)

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
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	body := make(map[string]interface{})
	body["agentid"] = d.oauth2.GetAgentId()
	body["event"] = *event
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRquest(http.MethodPost, "/topapi/calendar/v2/event/update", query, body)
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
	var result DintalkSuccessResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = xerrors.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = xerrors.Errorf("%s", "success is false")
		return
	}
	return
}

func (d *inner) CancelEvent(eventId string) (err error) {
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	body := make(map[string]interface{})
	body["event_id"] = eventId
	body["calendar_id"] = "primary"
	body["agentid"] = d.oauth2.GetAgentId()
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRquest(http.MethodPost, "/topapi/calendar/v2/event/cancel", query, body)
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
	var result DintalkSuccessResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = xerrors.Errorf("%s", result.Errmsg)
		return
	}
	if !result.Success {
		err = xerrors.Errorf("%s", "success is false")
		return
	}
	return
}
