package jpcal

import "time"

const dateFmt = "2006-01-02"

type Day interface {

	// Get date string
	// (e.g. "2006-01-01")
	Str() string

	// Get time.Time instance
	Time() (time.Time, error)

	//Get DayType
	Type() DayType

	//Get description of day in Japanese
	Description() string
}

type normalDay struct {
	date    string
	dayType DayType
}

type DayType string

const (
	TypeWeekDay         DayType = "weekday"
	TypeSaturday        DayType = "saturday"
	TypeSunday          DayType = "sunday"
	TypeNationalHoliday DayType = "national_holiday"
)

func (d *normalDay) Str() string {
	return d.date
}

func (d *normalDay) Time() (time.Time, error) {
	return time.Parse(dateFmt, d.date)
}

func (d *normalDay) Type() DayType {
	return d.dayType
}

func (d *normalDay) Description() string {
	switch d.dayType {
	case TypeWeekDay:
		return "平日"
	case TypeSaturday:
		return "土曜日"
	case TypeSunday:
		return "日曜日"
	}
	return ""
}

type nationalHoliday struct {
	date        string
	holidayName string
}

func (d *nationalHoliday) Str() string {
	return d.date
}

func (d *nationalHoliday) Time() (time.Time, error) {
	return time.Parse(dateFmt, d.date)
}

func (d *nationalHoliday) Type() DayType {
	return TypeNationalHoliday
}

func (d *nationalHoliday) Description() string {
	return d.holidayName
}
