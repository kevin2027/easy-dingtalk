package dingtalk

import (
	"github.com/kevin2027/easy-dingtalk/calendar"
	calendar_v2 "github.com/kevin2027/easy-dingtalk/calendar/v2"
	"github.com/kevin2027/easy-dingtalk/contact"
	"github.com/kevin2027/easy-dingtalk/message"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
)

type Dingtalk interface {
	SetDingDiReduceFn(fn utils.DingIdReduceFn)
	Oauth2() oauth2.Oauth2
	Contact() contact.Contact
	CalendarV2() calendar_v2.Calendar
	Calendar() calendar.Calendar
	Message() message.Message
}
