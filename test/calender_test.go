package dingtalk_test

import (
	"fmt"
	"testing"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	calendar_v2 "github.com/kevin2027/easy-dingtalk/calendar/v2"
)

func TestCalenderV2CreateEvent(t *testing.T) {
	var err error
	defer deferErr(&err)
	req := &calendar_v2.CreateEventRequestEvent{
		Attendees: []*calendar_v2.Attendee{
			{
				Userid: tea.String("user0"),
			},
		},
		CalendarId:  "",
		Description: tea.String("测试创建日程描述"),
		End: calendar_v2.DataTime{
			Timestamp: tea.Int64(time.Date(2022, 5, 2, 14, 0, 0, 0, time.Local).Unix()),
			Timezone:  tea.String("Asia/Shanghai"),
		},
		Start: calendar_v2.DataTime{
			Timestamp: tea.Int64(time.Date(2022, 5, 2, 13, 0, 0, 0, time.Local).Unix()),
			Timezone:  tea.String("Asia/Shanghai"),
		},
		Organizer: calendar_v2.Attendee{
			Userid: tea.String("user0"),
		},
		Summary:  "测试创建日程",
		Reminder: nil,
		Location: &calendar_v2.Location{
			Place: tea.String("地点"),
		},
		NotificationType: "",
	}
	res, err := client.CalendarV2().CreateEvent(req)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", *util.ToJSONString(res))
}

func TestCalendarV2UpdateEvent(t *testing.T) {
	var err error
	defer deferErr(&err)
	req := &calendar_v2.UpdateEventRequestEvent{
		Attendees:   []*calendar_v2.Attendee{},
		CalendarId:  "",
		Description: "测试修改日程描述",
		Start:       calendar_v2.DataTime{Timestamp: tea.Int64(time.Date(2022, 5, 2, 15, 0, 0, 0, time.Local).Unix()), Timezone: tea.String("Asia/Shanghai")},
		End:         calendar_v2.DataTime{Timestamp: tea.Int64(time.Date(2022, 5, 2, 16, 0, 0, 0, time.Local).Unix()), Timezone: tea.String("Asia/Shanghai")},
		Summary:     "测试修改日程",
		EventId:     "9E7066D46163091754634D654103262E",
		Reminder: &calendar_v2.Reminder{
			Method:  tea.String("app"),
			Minutes: tea.Int(5),
		},
		Location:  &calendar_v2.Location{Place: tea.String("地点")},
		Organizer: calendar_v2.Attendee{Userid: tea.String("user0")},
	}
	err = client.CalendarV2().UpdateEvent(req)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", "success")
}

func TestCalendarV2CancelEvent(t *testing.T) {
	var err error
	defer deferErr(&err)
	err = client.CalendarV2().CancelEvent("9E7066D46163091754634D654103262E")
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", "success")
}

func TestCalendarV2AttendeeUpdate(t *testing.T) {
	var err error
	defer deferErr(&err)
	attendeeList := []*calendar_v2.Attendee{
		{
			Userid:         tea.String("user1"),
			AttendeeStatus: tea.String("remove"),
		},
	}
	err = client.CalendarV2().AttendeeUpdate("9E7066D46163091754634D654103262E", attendeeList)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", "success")
}
