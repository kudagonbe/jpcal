package jpcal

import "time"

const dateFmt = "2006-01-02"

type Day interface {
	Time() (time.Time, error)
	Type() string
	Description() string
}

type NormalDay struct {
	date    string
	dayType dayType
}

type dayType string

const (
	TypeWeekDay         dayType = "weekday"
	TypeSaturday        dayType = "satuaday"
	TypeSunday          dayType = "sunday"
	TypeNationalHoliday dayType = "national_holiday"
)

func (d *NormalDay) Time() (time.Time, error) {
	return time.Parse(dateFmt, d.date)
}

func (d *NormalDay) Type() string {
	return string(d.dayType)
}

func (d *NormalDay) Description() string {
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

type NationalHoliday struct {
	date        string
	holidayName string
}

func (d *NationalHoliday) Time() (time.Time, error) {
	return time.Parse(dateFmt, d.date)
}

func (d *NationalHoliday) Type() string {
	return string(TypeNationalHoliday)
}

func (d *NationalHoliday) Description() string {
	return d.holidayName
}
