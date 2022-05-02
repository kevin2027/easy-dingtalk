package dingtalk_test

import (
	"fmt"
	"testing"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	calendar_v2 "github.com/kevin2027/easy-dingtalk/calendar/v2"

	"golang.org/x/xerrors"
)

func TestCreateCalenderV2(t *testing.T) {
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
		err = xerrors.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", *util.ToJSONString(res))
}
