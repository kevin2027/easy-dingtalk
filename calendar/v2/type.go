package calendar_v2

import "github.com/kevin2027/easy-dingtalk/utils"

type DataTime struct {
	Date      *string `json:"date,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty"`
	Timezone  *string `json:"timezone,omitempty"`
}

type DintalkSuccessResponse struct {
	utils.DintalkResponse
	Success bool `json:"success"`
}

type Location struct {
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Place     *string `json:"place,omitempty"`
}

type Reminder struct {
	Method  *string `json:"method,omitempty"`
	Minutes *int    `json:"minutes,omitempty"`
}

type Attendee struct {
	Userid         *string `json:"userid,omitempty"`
	AttendeeStatus *string `json:"attendee_status,omitempty"`
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
