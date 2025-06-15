# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`jpcal` is a Go library for acquiring and judging Japanese national holidays. It provides functionality to:
- Determine if a specific date is a weekday, Saturday, Sunday, or national holiday
- Get collections of days by type (weekdays, weekends, holidays) for a year or month
- Access Japanese holiday information from 1955-2026 based on Cabinet Office data

## Core Architecture

The library uses a type-based day classification system:
- `Day` interface with two implementations: `normalDay` and `nationalHoliday`
- `Days` type (slice of Day) with sorting capabilities
- Pre-generated calendar data in `calendar_gen.go` containing holiday mappings
- Four main day types: `TypeWeekDay`, `TypeSaturday`, `TypeSunday`, `TypeNationalHoliday`

Key files:
- `jpcal.go`: Main API functions for day type checking (`IsWeekday`, `IsSaturday`, etc.)
- `day.go`: Day interface and implementations
- `days.go`: Days collection type and bulk operations (`AllDays`, `NationalHolidays`, etc.)
- `calendar_gen.go`: Generated calendar data (DO NOT EDIT manually)
- `gen/gencalendar.go`: Code generator that fetches holiday data from Cabinet Office

## Development Commands

**Build:**
```bash
go build ./...
```

**Test:**
```bash
go test ./...
go test ./... -race -coverprofile=coverage.out -covermode=atomic
```

**Format:**
```bash
gofmt -l -w -s [filename]
```

**Generate calendar data:**
```bash
cd gen && go generate
```

## Code Generation

The calendar data is auto-generated from the Cabinet Office CSV. To update:
1. Run `go generate` in the `gen/` directory
2. This fetches fresh data and regenerates `calendar_gen.go`
3. The generator includes `//go:generate` directives for formatting

## Testing

All main functionality has corresponding test files (`*_test.go`). Tests cover:
- Individual day operations
- Bulk day collections
- Edge cases and error conditions
- Holiday detection accuracy