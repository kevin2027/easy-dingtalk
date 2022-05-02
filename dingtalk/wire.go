//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package dingtalk

import (
	"github.com/Kevin2027/easy-dingtalk/calendar"
	calendar_v2 "github.com/Kevin2027/easy-dingtalk/calendar/v2"
	"github.com/Kevin2027/easy-dingtalk/contact"
	"github.com/Kevin2027/easy-dingtalk/oauth2"
	"github.com/Kevin2027/easy-dingtalk/utils"
	"github.com/google/wire"
)

func NewDingtalk(opt utils.DingtalkOptions) (Dingtalk, func(), error) {
	panic(wire.Build(calendar_v2.NewCalendar, calendar.NewCalendar, oauth2.NewOuath2, contact.NewContact, newDingtalk))
}
