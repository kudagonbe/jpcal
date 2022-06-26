package jpcal

import "time"

const dateFmt = "2006-01-02"

type Day interface {
	Str() string
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
	TypeSaturday        dayType = "saturday"
	TypeSunday          dayType = "sunday"
	TypeNationalHoliday dayType = "national_holiday"
)

func (d *NormalDay) Str() string {
	return d.date
}

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

func (d *NationalHoliday) Str() string {
	return d.date
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

type Days []Day

func (ds Days) Len() int {
	return len(ds)
}

func (ds Days) Less(i, j int) bool {
	return ds[i].Str() < ds[j].Str()
}

func (ds Days) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}
