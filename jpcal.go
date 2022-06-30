package jpcal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Judge if the specified day is a weekday.
func IsWeekday(year int, month int, day int) (bool, error) {
	return isSpecificTypeDay(year, month, day, TypeWeekDay)
}

// Judge if the specified day is a saturday.
// (except for national holiday)
func IsSaturday(year int, month int, day int) (bool, error) {
	return isSpecificTypeDay(year, month, day, TypeSaturday)
}

// Judge if the specified day is a sunday.
// (except for national holiday)
func IsSunday(year int, month int, day int) (bool, error) {
	return isSpecificTypeDay(year, month, day, TypeSunday)
}

// Judge if the specified day is a national holiday.
func IsNationalHoliday(year int, month int, day int) (bool, error) {
	return isSpecificTypeDay(year, month, day, TypeNationalHoliday)
}

func isSpecificTypeDay(year int, month int, day int, dt DayType) (bool, error) {
	if err := chkYear(year); err != nil {
		return false, err
	}
	if err := chkMonth(month); err != nil {
		return false, err
	}

	var m map[int]map[int]string
	var dayStr string = strconv.Itoa(day)

	switch dt {
	case TypeWeekDay:
		m = weekdays
	case TypeSaturday:
		m = saturdays
	case TypeSunday:
		m = sundays
	case TypeNationalHoliday:
		m = holidays
	default:
		return false, errors.New("invalid day type")
	}

	for _, v := range strings.Split(m[year][month], ",") {
		if v == dayStr {
			return true, nil
		}
	}
	return false, nil
}

func chkYear(year int) error {
	if year < MinYear {
		return fmt.Errorf("jpcal is only supported after %d, but the year you set is %d", MinYear, year)
	}

	if year > MaxYear {
		return fmt.Errorf("jpcal is only supported until %d, but the year you set is %d", MaxYear, year)
	}

	return nil
}

func chkMonth(month int) error {
	if month < int(time.January) || int(time.December) < month {
		return fmt.Errorf("invalid month: %d", month)
	}

	return nil
}
