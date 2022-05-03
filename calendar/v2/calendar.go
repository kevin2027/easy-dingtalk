package calendar_v2

import (
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type Calendar interface {
	CreateEvent(event *CreateEventRequestEvent) (res *CreateEventResponseResult, err error)
	UpdateEvent(event *UpdateEventRequestEvent) (err error)
	CancelEvent(eventId string) (err error)
	AttendeeUpdate(eventId string, eventAttendees []*Attendee) (err error)
	utils.DingIdReduceAble
}

func NewCalendar(oauth2 oauth2.Oauth2) (r Calendar) {
	return &inner{
		oauth2: oauth2,
	}
}
