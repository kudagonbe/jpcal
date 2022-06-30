package jpcal

import (
	"fmt"
	"time"
)

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
	year    int
	month   int
	day     int
	dayType DayType
}

type DayType string

const (
	TypeWeekDay         DayType = "weekday"  // except for national holiday
	TypeSaturday        DayType = "saturday" // except for national holiday
	TypeSunday          DayType = "sunday"   // except for national holiday
	TypeNationalHoliday DayType = "national_holiday"
)

func (d *normalDay) Str() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.year, d.month, d.day)
}

func (d *normalDay) Time() (time.Time, error) {
	return time.Date(d.year, time.Month(d.month), d.day, 0, 0, 0, 0, time.UTC), nil
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
	year        int
	month       int
	day         int
	holidayName string
}

func (d *nationalHoliday) Str() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.year, d.month, d.day)
}

func (d *nationalHoliday) Time() (time.Time, error) {
	return time.Date(d.year, time.Month(d.month), d.day, 0, 0, 0, 0, time.UTC), nil
}

func (d *nationalHoliday) Type() DayType {
	return TypeNationalHoliday
}

func (d *nationalHoliday) Description() string {
	return d.holidayName
}

func GetDay(year int, month int, day int) (Day, error) {
	ds, err := AllDaysYM(year, month)
	if err != nil {
		return nil, err
	}

	dayStr := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	for _, d := range ds {
		if dayStr == d.Str() {
			return d, nil
		}
	}

	return nil, fmt.Errorf("invalid argument::year:%d month:%d day:%d", year, month, day)
}
