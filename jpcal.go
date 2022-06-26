package jpcal

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Holidays(year int) (Days, error) {
	var ds Days = make([]Day, 0, 366)

	if err := chkYear(year); err != nil {
		return nil, err
	}

	if newds, err := appendHolidays(ds, year); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	sort.Sort(ds)
	return ds, nil
}

func HolidaysYM(year int, month int) (Days, error) {
	var ds Days = make([]Day, 0, 31)

	if err := chkYear(year); err != nil {
		return nil, err
	}
	if err := chkMonth(month); err != nil {
		return nil, err
	}

	if newds, err := appendHolidaysYM(ds, year, month); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	sort.Sort(ds)
	return ds, nil
}

func AllDays(year int) (Days, error) {
	var ds Days = make([]Day, 0, 366)
	if err := chkYear(year); err != nil {
		return nil, err
	}

	if newds, err := appendWeekdays(ds, year); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	if newds, err := appendSaturdays(ds, year); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	if newds, err := appendSundays(ds, year); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	if newds, err := appendHolidays(ds, year); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	sort.Sort(ds)
	return ds, nil
}

func AllDaysYM(year int, month int) (Days, error) {
	var ds Days = make([]Day, 0, 31)

	if err := chkYear(year); err != nil {
		return nil, err
	}

	if err := chkMonth(month); err != nil {
		return nil, err
	}

	if newds, err := appendWeekdaysYM(ds, year, month); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	if newds, err := appendSaturdaysYM(ds, year, month); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	if newds, err := appendSundaysYM(ds, year, month); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	if newds, err := appendHolidaysYM(ds, year, month); err != nil {
		return nil, err
	} else {
		ds = newds
	}

	sort.Sort(ds)
	return ds, nil
}

func SpecificTypeDays(year int, ts ...dayType) (Days, error) {
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
		if newds, err := appendWeekdays(ds, year); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sat {
		if newds, err := appendSaturdays(ds, year); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sun {
		if newds, err := appendSundays(ds, year); err != nil {
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

func SpecificTypeDaysYM(year int, month int, ts ...dayType) (Days, error) {
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
		if newds, err := appendWeekdaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sat {
		if newds, err := appendSaturdaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}

	if sun {
		if newds, err := appendSundaysYM(ds, year, month); err != nil {
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

func appendHolidays(ds Days, year int) (Days, error) {
	if _, ok := holidays[year]; !ok {
		return ds, nil
	}

	for month := range holidays[year] {
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
		ds = append(ds, &NationalHoliday{
			date:        fmt.Sprintf("%04d-%02d-%02d", year, month, d),
			holidayName: hns[i],
		})
	}
	return ds, nil
}

func appendWeekdays(ds Days, year int) (Days, error) {
	if _, ok := weekdays[year]; !ok {
		return ds, nil
	}

	for month := range weekdays[year] {
		if newds, err := appendWeekdaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}
	return ds, nil
}

func appendWeekdaysYM(ds Days, year int, month int) (Days, error) {
	if _, ok := weekdays[year]; !ok {
		return ds, nil
	}

	if _, ok := weekdays[year][month]; !ok {
		return ds, nil
	}

	hs := strings.Split(weekdays[year][month], ",")

	for _, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ds = append(ds, &NormalDay{
			date:    fmt.Sprintf("%04d-%02d-%02d", year, month, d),
			dayType: TypeWeekDay,
		})
	}
	return ds, nil
}

func appendSaturdays(ds Days, year int) (Days, error) {
	if _, ok := saturdays[year]; !ok {
		return ds, nil
	}

	for month := range saturdays[year] {
		if newds, err := appendSaturdaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}
	return ds, nil
}

func appendSaturdaysYM(ds Days, year int, month int) (Days, error) {
	if _, ok := saturdays[year]; !ok {
		return ds, nil
	}

	if _, ok := saturdays[year][month]; !ok {
		return ds, nil
	}

	hs := strings.Split(saturdays[year][month], ",")

	for _, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ds = append(ds, &NormalDay{
			date:    fmt.Sprintf("%04d-%02d-%02d", year, month, d),
			dayType: TypeSaturday,
		})
	}
	return ds, nil
}

func appendSundays(ds Days, year int) (Days, error) {
	if _, ok := sundays[year]; !ok {
		return ds, nil
	}

	for month := range sundays[year] {
		if newds, err := appendSundaysYM(ds, year, month); err != nil {
			return nil, err
		} else {
			ds = newds
		}
	}
	return ds, nil
}

func appendSundaysYM(ds Days, year int, month int) (Days, error) {
	if _, ok := sundays[year]; !ok {
		return ds, nil
	}

	if _, ok := sundays[year][month]; !ok {
		return ds, nil
	}

	hs := strings.Split(sundays[year][month], ",")

	for _, v := range hs {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ds = append(ds, &NormalDay{
			date:    fmt.Sprintf("%04d-%02d-%02d", year, month, d),
			dayType: TypeSunday,
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
