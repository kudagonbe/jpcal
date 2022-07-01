package jpcal

import (
	"sort"
	"strconv"
	"strings"
)

type Days []Day

// for override sort.Interface
func (ds Days) Len() int {
	return len(ds)
}

// for override sort.Interface
func (ds Days) Less(i, j int) bool {
	return ds[i].Str() < ds[j].Str()
}

// for override sort.Interface
func (ds Days) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

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
		ds = appendNormalDays(ds, year, TypeWeekDay)
	}

	if sat {
		ds = appendNormalDays(ds, year, TypeSaturday)
	}

	if sun {
		ds = appendNormalDays(ds, year, TypeSunday)
	}

	if hd {
		ds = appendHolidays(ds, year)
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
		ds = appendNormalDaysYM(ds, year, month, TypeWeekDay)
	}

	if sat {
		ds = appendNormalDaysYM(ds, year, month, TypeSaturday)
	}

	if sun {
		ds = appendNormalDaysYM(ds, year, month, TypeSunday)
	}

	if hd {
		ds = appendHolidaysYM(ds, year, month)
	}

	sort.Sort(ds)
	return ds, nil
}

func appendNormalDays(ds Days, year int, dt DayType) Days {
	for month := 1; month <= 12; month++ {
		ds = appendNormalDaysYM(ds, year, month, dt)
	}
	return ds
}

func appendNormalDaysYM(ds Days, year int, month int, dt DayType) Days {
	var m map[int]map[int]string
	switch dt {
	case TypeWeekDay:
		m = weekdays
	case TypeSaturday:
		m = saturdays
	case TypeSunday:
		m = sundays
	default:
		return ds
	}
	if _, ok := m[year]; !ok {
		return ds
	}

	if _, ok := m[year][month]; !ok {
		return ds
	}

	hs := strings.Split(m[year][month], ",")

	for _, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return ds
		}
		ds = append(ds, &normalDay{
			year:    year,
			month:   month,
			day:     d,
			dayType: dt,
		})
	}
	return ds
}

func appendHolidays(ds Days, year int) Days {
	for month := 1; month <= 12; month++ {
		ds = appendHolidaysYM(ds, year, month)
	}
	return ds
}

func appendHolidaysYM(ds Days, year int, month int) Days {
	if _, ok := holidays[year]; !ok {
		return ds
	}

	if _, ok := holidays[year][month]; !ok {
		return ds
	}

	hs := strings.Split(holidays[year][month], ",")
	hns := strings.Split(holidayNames[year][month], ",")

	for i, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return ds
		}
		ds = append(ds, &nationalHoliday{
			year:        year,
			month:       month,
			day:         d,
			holidayName: hns[i],
		})
	}
	return ds
}
