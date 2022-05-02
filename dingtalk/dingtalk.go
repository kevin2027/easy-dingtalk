package dingtalk

import (
	"github.com/Kevin2027/easy-dingtalk/calendar"
	calendar_v2 "github.com/Kevin2027/easy-dingtalk/calendar/v2"
	"github.com/Kevin2027/easy-dingtalk/contact"
	"github.com/Kevin2027/easy-dingtalk/oauth2"
	"github.com/Kevin2027/easy-dingtalk/utils"
)

type Dingtalk interface {
	SetDingDiReduceFn(fn utils.DingIdReduceFn)
	Oauth2() oauth2.Oauth2
	Contact() contact.Contact
	CalendarV2() calendar_v2.Calendar
	Calendar() calendar.Calendar
}

type inner struct {
	oauth2     oauth2.Oauth2
	contact    contact.Contact
	calendarV2 calendar_v2.Calendar
	calendar   calendar.Calendar
}

func newDingtalk(oauth2 oauth2.Oauth2,
	contact contact.Contact,
	calendarV2 calendar_v2.Calendar,
	calendar calendar.Calendar) Dingtalk {
	return &inner{
		oauth2:     oauth2,
		calendarV2: calendarV2,
		calendar:   calendar,
		contact:    contact,
	}
}

func (d *inner) SetDingDiReduceFn(fn utils.DingIdReduceFn) {
	d.calendar.SetDingDiReduceFn(fn)
	d.calendarV2.SetDingDiReduceFn(fn)
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
