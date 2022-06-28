package jpcal

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Get all national holidays in a year.
func NationalHolidays(year int) (Days, error) {
	return SpecificTypeDays(year, TypeNationalHoliday)
}

// Get all national holidays in a month.
func NationalHolidaysYM(year int, month int) (Days, error) {
	return SpecificTypeDaysYM(year, month, TypeNationalHoliday)
}

// Get all days in a year.
func AllDays(year int) (Days, error) {
	return SpecificTypeDays(year, TypeWeekDay, TypeSaturday, TypeSunday, TypeNationalHoliday)
}

// Get all days in a month.
func AllDaysYM(year int, month int) (Days, error) {
	return SpecificTypeDaysYM(year, month, TypeWeekDay, TypeSaturday, TypeSunday, TypeNationalHoliday)
}

// Get specific type days in a year.
// You can choose day type 'TypeWeekDay', 'TypeSaturday', 'TypeSunday', 'TypeNationalHoliday'
func SpecificTypeDays(year int, ts ...DayType) (Days, error) {
	var ds Days = make([]Day, 0, 366)
	var wd, sat, sun, hd bool

	if err := chkYear(year); err != nil {
		return nil, err
	}

	for _, v := range ts {
		switch v {
		case TypeWeekDay:
			wd = true
		case TypeSaturday:
			sat = true
		case TypeSunday:
			sun = true
		case TypeNationalHoliday:
			hd = true
		}
	}

	if wd {
		if newds, err := appendNormalDays(ds, year, TypeWeekDay); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sat {
		if newds, err := appendNormalDays(ds, year, TypeSaturday); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sun {
		if newds, err := appendNormalDays(ds, year, TypeSunday); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if hd {
		if newds, err := appendHolidays(ds, year); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	sort.Sort(ds)
	return ds, nil
}

// Get specific type days in a month.
// You can choose day type 'TypeWeekDay', 'TypeSaturday', 'TypeSunday', 'TypeNationalHoliday'
func SpecificTypeDaysYM(year int, month int, ts ...DayType) (Days, error) {
	var ds Days = make([]Day, 0, 31)
	var wd, sat, sun, hd bool

	if err := chkYear(year); err != nil {
		return nil, err
	}

	if err := chkMonth(month); err != nil {
		return nil, err
	}

	for _, v := range ts {
		switch v {
		case TypeWeekDay:
			wd = true
		case TypeSaturday:
			sat = true
		case TypeSunday:
			sun = true
		case TypeNationalHoliday:
			hd = true
		}
	}

	if wd {
		if newds, err := appendNormalDaysYM(ds, year, month, TypeWeekDay); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sat {
		if newds, err := appendNormalDaysYM(ds, year, month, TypeSaturday); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sun {
		if newds, err := appendNormalDaysYM(ds, year, month, TypeSunday); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if hd {
		if newds, err := appendHolidaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	sort.Sort(ds)
	return ds, nil
}

func appendNormalDays(ds Days, year int, dt DayType) (Days, error) {
	for month := 1; month <= 12; month++ {
		if newds, err := appendNormalDaysYM(ds, year, month, dt); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}
	return ds, nil
}

func appendNormalDaysYM(ds Days, year int, month int, dt DayType) (Days, error) {
	var m map[int]map[int]string
	switch dt {
	case TypeWeekDay:
		m = weekdays
	case TypeSaturday:
		m = saturdays
	case TypeSunday:
		m = sundays
	default:
		return nil, errors.New("invalid day type")
	}
	if _, ok := m[year]; !ok {
		return ds, nil
	}

	if _, ok := m[year][month]; !ok {
		return ds, nil
	}

	hs := strings.Split(m[year][month], ",")

	for _, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ds = append(ds, &normalDay{
			date:    fmt.Sprintf("%04d-%02d-%02d", year, month, d),
			dayType: dt,
		})
	}
	return ds, nil
}

func appendHolidays(ds Days, year int) (Days, error) {
	for month := 1; month <= 12; month++ {
		if newds, err := appendHolidaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}
	return ds, nil
}

func appendHolidaysYM(ds Days, year int, month int) (Days, error) {
	if _, ok := holidays[year]; !ok {
		return ds, nil
	}

	if _, ok := holidays[year][month]; !ok {
		return ds, nil
	}

	hs := strings.Split(holidays[year][month], ",")
	hns := strings.Split(holidayNames[year][month], ",")

	for i, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ds = append(ds, &nationalHoliday{
			date:        fmt.Sprintf("%04d-%02d-%02d", year, month, d),
			holidayName: hns[i],
		})
	}
	return ds, nil
}

func chkYear(year int) error {
	if year < minYear {
		return fmt.Errorf("jpcal is only supported after %d, but the year you set is %d", minYear, year)
	}

	if year > maxYear {
		return fmt.Errorf("jpcal is only supported until %d, but the year you set is %d", maxYear, year)
	}

	return nil
}

func chkMonth(month int) error {
	if month < int(time.January) || int(time.December) < month {
		return fmt.Errorf("invalid month: %d", month)
	}

	return nil
}

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
