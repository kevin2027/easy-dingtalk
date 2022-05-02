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
