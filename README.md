# jpcal

[![Go Reference](https://pkg.go.dev/badge/github.com/kudagonbe/jpcal.svg)](https://pkg.go.dev/github.com/kudagonbe/jpcal)
[![CI Job](https://github.com/kudagonbe/jpcal/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kudagonbe/jpcal/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kudagonbe/jpcal/branch/main/graph/badge.svg?token=E3CJKEO0D5)](https://codecov.io/gh/kudagonbe/jpcal)

This is a library made in Go language for acquiring and judging Japanese holidays.

## Installing

First, use go get to install the latest version of the library.

```
go get -u github.com/kudagonbe/jpcal@latest
```

Next, include jpcal in your application:

```
import "github.com/kudagonbe/jpcal"
```

## Example

1. Get `Day` instance

```go
var day Day

day, _ := jpcal.GetDay(2020, 1, 1)

str := day.Str() // "2020-01-01"
t := day.Time() // instance of time.Time
dt := day.Type() // jpcal.TypeNationalHoliday
desc := day.Description() // "元日"
```

2. Get `Days` instance

You can use public functions that return `Days` instance.
(type `Days` expand type `[]Day`)
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
ds, err = jpcal.NationalHolidays(year)

// Get all national holidays in a month.
ds, err = jpcal.NationalHolidaysYM(year, month)

// Get days of specific type in a year.
ds, err = jpcal.SpecificTypeDays(year, jpcal.TypeWeekday, jpcal.TypeSaturday)

// Get days of specific type in a month.
ds, err = jpcal.SpecificTypeDaysYM(year, month, jpcal.TypeSunday, jpcal.TypeNationalHoliday)
```

3. Judge DayType

You can judge DayType.

```go
// Judge if the specified day is a weekday(except for national holiday).
result, err := jpcal.IsWeekday(2020, 1, 2) //true, nil

// Judge if the specified day is a saturday(except for national holiday).
result, err := jpcal.IsSaturday(2020, 1, 4) //true, nil

// Judge if the specified day is a sunday(except for national holiday).
result, err := jpcal.IsSunday(2020, 1, 5) //true, nil

// Judge if the specified day is a national holiday.
result, err := jpcal.IsNationalHoliday(2020, 1, 1) //true, nil

```