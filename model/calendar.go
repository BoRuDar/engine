package model

import (
	"encoding/json"
	"io"
	"net/http"
)

type CalendarAcceptOfDay struct {
	Id             int64 `json:"id" db:"id"`
	Week           int8  `json:"week" db:"week_day"`
	StartTimeOfDay int16 `json:"start_time_of_day" db:"start_time_of_day"`
	EndTimeOfDay   int16 `json:"end_time_of_day" db:"end_time_of_day"`
}

type CalendarExceptDate struct {
	Id         int64 `json:"id"`
	CalendarId int64 `json:"calendar_id"`
	Repeat     int8  `json:"repeat"`
	Date       int   `json:"date"`
}

// Description of the Calendar
// swagger:model Calendar
type Calendar struct {
	DomainRecord
	Name        string  `json:"name" db:"name"`
	Start       *int64  `json:"start" db:"start"`
	Finish      *int64  `json:"finish" db:"finish"`
	Timezone    Lookup  `json:"timezone"`
	Description *string `json:"description,omitempty"`
}

type Timezone struct {
	Id     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Offset string `json:"offset" db:"offset"`
}

func (c *Calendar) IsValid() *AppError {
	if len(c.Name) <= 3 {
		return NewAppError("Calendar.IsValid", "model.calendar.is_valid.name.app_error", nil, "name="+c.Name, http.StatusBadRequest)
	}

	if c.DomainId == 0 {
		return NewAppError("Calendar.IsValid", "model.calendar.is_valid.domain_id.app_error", nil, "name="+c.Name, http.StatusBadRequest)
	}
	return nil
}

func CalendarFromJson(data io.Reader) *Calendar {
	var calendar Calendar
	if err := json.NewDecoder(data).Decode(&calendar); err != nil {
		return nil
	} else {
		return &calendar
	}
}

func CalendarsToJson(calendars []*Calendar) string {
	b, _ := json.Marshal(calendars)
	return string(b)
}

func (c *Calendar) ToJson() string {
	b, _ := json.Marshal(c)
	return string(b)
}
