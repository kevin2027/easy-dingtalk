package dingtalk

import (
	"github.com/kevin2027/easy-dingtalk/calendar"
	calendar_v2 "github.com/kevin2027/easy-dingtalk/calendar/v2"
	"github.com/kevin2027/easy-dingtalk/contact"
	"github.com/kevin2027/easy-dingtalk/meeting"
	"github.com/kevin2027/easy-dingtalk/message"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type inner struct {
	oauth2     oauth2.Oauth2
	contact    contact.Contact
	calendarV2 calendar_v2.Calendar
	calendar   calendar.Calendar
	message    message.Message
	meeting    meeting.Meeting
}

func newDingtalk(oauth2 oauth2.Oauth2,
	contact contact.Contact,
	calendarV2 calendar_v2.Calendar,
	calendar calendar.Calendar, message message.Message, meeting meeting.Meeting) Dingtalk {
	return &inner{
		oauth2:     oauth2,
		calendarV2: calendarV2,
		calendar:   calendar,
		contact:    contact,
		message:    message,
		meeting:    meeting,
	}
}

func (d *inner) SetDingDiReduceFn(fn utils.DingIdReduceFn) {
	d.calendar.SetReduceFn(fn)
	d.calendarV2.SetReduceFn(fn)
	d.message.SetReduceFn(fn)
	d.contact.SetReduceFn(fn)
	d.meeting.SetReduceFn(fn)
}

func (d *inner) Oauth2() oauth2.Oauth2 {
	return d.oauth2
}

func (d *inner) Contact() contact.Contact {
	return d.contact
}
func (d *inner) CalendarV2() calendar_v2.Calendar {
	return d.calendarV2
}
func (d *inner) Calendar() calendar.Calendar {
	return d.calendar
}

func (d *inner) Message() message.Message {
	return d.message
}

func (d *inner) Meeting() meeting.Meeting {
	return d.meeting
}
