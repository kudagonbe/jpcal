# jp-calendar

## Description

This is a library made in Go language for acquiring and judging Japanese holidays.

日本の平日・土日・祝日を取得・判定するためのGo言語製のライブラリです。

## Usage

### Get **Days** instance

You can use public functions that return `Days` instance.

```go
	const year int = 2020
	const month int = 1

	var ds jpcal.Days
    var err error

	// Get all days in a year.
	ds, err = jpcal.AllDays(year)

	// Get all days in a month.
	ds, err = jpcal.AllDaysYM(year, month)

	// Get all national holidays in a year.
	ds, err = jpcal.Holidays(year)

	// Get all national holidays in a month.
	ds, err = jpcal.HolidaysYM(year, month)

	// Get specific type days in a year.
    // You can choose day type 'TypeWeekDay', 'TypeSaturday', 'TypeSunday', 'TypeNationalHoliday'
	ds, err = jpcal.SpecificTypeDays(year, jpcal.TypeSaturday, jpcal.TypeSunday)

	// Get specific type days in a month.
    // You can choose day type 'TypeWeekDay', 'TypeSaturday', 'TypeSunday', 'TypeNationalHoliday'
	ds, err = jpcal.SpecificTypeDaysYM(year, month, jpcal.TypeWeekDay, jpcal.TypeNationalHoliday)
```

### Usage of type **Day**

`Days` is expanded type of `[]Day`.
So you should know usage of type `Day`.

```go
type Day interface {
    /*
    Get date string
    e.g. "2006-01-01"
    */
    Str() string

    /*
    Get time.Time instance
    */
	Time() (time.Time, error)

    /*
    Get type of day
    "weekday", "saturday", "sunday" and "national_holiday"
    */
	Type() string

    /*
    Get description of day in Japanese
    */
	Description() string
}
```

