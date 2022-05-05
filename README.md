# dingtalk

这是一个集成钉钉常用功能的简易版服务端开发工具库,适用于创建一次客户端，多次调用的场景。内置token过期时间维护；内置用户id到用户userid的计算函数，可以方便的在企业用户名与userid之间自动转换；同时在reduce函数中可以设置加入白名单过滤，避免在测试环境中发信息给非白名单用户。具体参数说明，请参考[钉钉开发文档](https://open.dingtalk.com/document/orgapp-server/api-overview)

## 安装

```shell
  go get -u github.com/kevin2027/easy-dingtalk
```

## 引入代码

```go
import (
    "github.com/kevin2027/easy-dingtalk/dingtalk"
    "github.com/kevin2027/easy-dingtalk/utils"
)
```

## 创建客户端

```go
srv, _, err = dingtalk.NewDingtalk(utils.DingtalkOptions{
    AppKey:    config.AppKey,
    AppSecret: config.AppSecret,
    AgentId:   config.AgentId,
})

```

## reduce函数

```go
client.SetDingDiReduceFn(func(ctx context.Context, attr string, src ...string) (dest map[string]string) {
  dest = make(map[string]string)
  if attr == utils.AttDeptId {
   return
  }
  for _, s := range src {
   if user, ok := config.Users[s]; ok {
    switch attr {
    case utils.AttrUserid:
     dest[s] = user.Userid
    }
   }
  }

  return
 })
```

## 调用

### 消息通知

#### 普通消息

```go
var err error
 defer deferErr(&err)
 msg := &message.MessageRequest{
  Msgtype: "text",
  Text: &message.TextMessage{
   Content: "这是一段文本消息",
  },
 }
 receiver, err := client.Message().SendToConversation("user0", 123453556, msg)
 if err != nil {
  err = fmt.Errorf("%w", err)
  return
 }
 fmt.Printf("%v\n", receiver)
```

#### 工作通知

```go
 var err error
 defer deferErr(&err)
 msg := &message.MessageRequest{
  Msgtype: "text",
  Text: &message.TextMessage{
   Content: "这是一段文本消息",
  },
 }
 taskId, err := client.Message().CorpconversationaSyncsendV2([]string{"user0"}, nil, false, msg)
 if err != nil {
  err = fmt.Errorf("%w", err)
  return
 }
 fmt.Printf("%v\n", taskId)

```

### 日程（新版）

#### 添加日程

```go
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

```

#### 修改日程

```go
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
```

#### 取消日程

```go
err = client.CalendarV2().CancelEvent("9E7066D46163091754634D654103262E")
 if err != nil {
  err = fmt.Errorf("%w", err)
  return
 }
 fmt.Printf("%v\n", "success")
```

#### 修改日程参与者

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

### 日程（旧版）

```go

 CreateEvent(unionId string, req *dingtalkcalendar_1_0.CreateEventRequest) (event *dingtalkcalendar_1_0.CreateEventResponseBody, err error)

 PatchEvent(unionId string, eventId string, req *dingtalkcalendar_1_0.PatchEventRequest) (event *dingtalkcalendar_1_0.PatchEventResponseBody, err error)

 DeleteEvent(unionId string, eventId string) (err error)

 AddAttendee(unionId string, eventId string, req *dingtalkcalendar_1_0.AddAttendeeRequest) (err error)

 RemoveAttendee(unionId string, eventId string, req *dingtalkcalendar_1_0.RemoveAttendeeRequest) (err error)
```

### 会议

```go
 CreateVideoConference(userId string, confTitle string, inviteUserIds []string, inviteCaller bool) (res *dingtalkconference_1_0.  CreateVideoConferenceResponseBody, err error)

 CloseVideoConference(unionId string, conferenceId string) (err error)

 QueryConferenceInfoBatch(conferenceIdList []string) (res []*dingtalkconference_1_0.QueryConferenceInfoBatchResponseBodyInfos, err error)
```

### 通讯录

```go
GetUserInfo(userid string) (res *GetUserInfoResponseResult, err error)
```

### Oauth2

```go
  GetAccessToken() (accessToken string, expireIn time.Time, err error)
  
  GetAgentId() (agentId int64)

  SetAgentId(agentId int64)

  GetUserToken(code string, refreshToken string) (res *dingtalkoauth2_1_0.GetUserTokenResponseBody, err error)
```
