# jpcal

[![Go Reference](https://pkg.go.dev/badge/github.com/kudagonbe/jpcal.svg)](https://pkg.go.dev/github.com/kudagonbe/jpcal)
[![CI Job](https://github.com/kudagonbe/jpcal/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kudagonbe/jpcal/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kudagonbe/jpcal/branch/main/graph/badge.svg?token=E3CJKEO0D5)](https://codecov.io/gh/kudagonbe/jpcal)

This is a library made in Go language for acquiring and judging Japanese holidays.

## Example

1. Get `Days` instance

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
ds, err = jpcal.NationalHolidays(year)

// Get all national holidays in a month.
ds, err = jpcal.NationalHolidaysYM(year, month)

```

2. Judge DayType

You can judge DayType.

```go
// Judge if the specified day is a national holiday.
result, err := jpcal.IsNationalHoliday(2020, 1, 1) //true, nil

```