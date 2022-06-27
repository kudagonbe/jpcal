package jpcal

import (
	"reflect"
	"testing"
)

var wantHolidays2020 Days = []Day{
	&NationalHoliday{date: "2020-01-01", holidayName: "元日"},
	&NationalHoliday{date: "2020-01-13", holidayName: "成人の日"},
	&NationalHoliday{date: "2020-02-11", holidayName: "建国記念の日"},
	&NationalHoliday{date: "2020-02-23", holidayName: "天皇誕生日"},
	&NationalHoliday{date: "2020-02-24", holidayName: "休日"},
	&NationalHoliday{date: "2020-03-20", holidayName: "春分の日"},
	&NationalHoliday{date: "2020-04-29", holidayName: "昭和の日"},
	&NationalHoliday{date: "2020-05-03", holidayName: "憲法記念日"},
	&NationalHoliday{date: "2020-05-04", holidayName: "みどりの日"},
	&NationalHoliday{date: "2020-05-05", holidayName: "こどもの日"},
	&NationalHoliday{date: "2020-05-06", holidayName: "休日"},
	&NationalHoliday{date: "2020-07-23", holidayName: "海の日"},
	&NationalHoliday{date: "2020-07-24", holidayName: "スポーツの日"},
	&NationalHoliday{date: "2020-08-10", holidayName: "山の日"},
	&NationalHoliday{date: "2020-09-21", holidayName: "敬老の日"},
	&NationalHoliday{date: "2020-09-22", holidayName: "秋分の日"},
	&NationalHoliday{date: "2020-11-03", holidayName: "文化の日"},
	&NationalHoliday{date: "2020-11-23", holidayName: "勤労感謝の日"},
}

func TestHolidays(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name    string
		args    args
		want    Days
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{year: 2020},
			want:    wantHolidays2020,
			wantErr: false,
		},
		{
			name:    "too_small_year",
			args:    args{year: minYear - 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: maxYear + 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Holidays(tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("Holidays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Holidays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHolidaysYM(t *testing.T) {
	type args struct {
		year  int
		month int
	}
	tests := []struct {
		name    string
		args    args
		want    Days
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{year: 2020, month: 1},
			want:    wantHolidays2020[0:2],
			wantErr: false,
		},
		{
			name:    "too_small_year",
			args:    args{year: minYear - 1, month: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: maxYear + 1, month: 12},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too_small_month",
			args:    args{year: 2020, month: 0},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too_big_month",
			args:    args{year: 2020, month: 13},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HolidaysYM(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("HolidaysYM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HolidaysYM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllDays(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "success",
			args:      args{year: 2020},
			wantCount: 366,
			wantErr:   false,
		},
		{
			name:      "too_small_year",
			args:      args{year: minYear - 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: maxYear + 1},
			wantCount: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AllDays(tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllDays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Len() != tt.wantCount {
				t.Errorf("ActualCount = %d, wantCount %v", got.Len(), tt.wantCount)
			}
		})
	}
}

func TestAllDaysYM(t *testing.T) {
	type args struct {
		year  int
		month int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "success",
			args:      args{year: 2020, month: 1},
			wantCount: 31,
			wantErr:   false,
		},
		{
			name:      "too_small_year",
			args:      args{year: minYear - 1, month: 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: maxYear + 1, month: 12},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_small_month",
			args:      args{year: 2020, month: 0},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_month",
			args:      args{year: 2020, month: 13},
			wantCount: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AllDaysYM(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllDaysYM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Len() != tt.wantCount {
				t.Errorf("ActualCount = %d, wantCount %v", got.Len(), tt.wantCount)
			}
		})
	}
}

func TestSpecificTypeDays(t *testing.T) {
	type args struct {
		year int
		ts   []dayType
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "success_weekday",
			args:      args{year: 2020, ts: []dayType{TypeWeekDay}},
			wantCount: 246,
			wantErr:   false,
		},
		{
			name:      "success_saturday",
			args:      args{year: 2020, ts: []dayType{TypeSaturday}},
			wantCount: 52,
			wantErr:   false,
		},
		{
			name:      "success_sunday",
			args:      args{year: 2020, ts: []dayType{TypeSunday}},
			wantCount: 50,
			wantErr:   false,
		},
		{
			name:      "success_holiday",
			args:      args{year: 2020, ts: []dayType{TypeNationalHoliday}},
			wantCount: 18,
			wantErr:   false,
		},
		{
			name:      "success_weekday_and_saturday",
			args:      args{year: 2020, ts: []dayType{TypeWeekDay, TypeSaturday}},
			wantCount: 298,
			wantErr:   false,
		},
		{
			name:      "success_sunday_and_holiday",
			args:      args{year: 2020, ts: []dayType{TypeSunday, TypeNationalHoliday}},
			wantCount: 68,
			wantErr:   false,
		},
		{
			name:      "success_all",
			args:      args{year: 2020, ts: []dayType{TypeWeekDay, TypeSaturday, TypeSunday, TypeNationalHoliday}},
			wantCount: 366,
			wantErr:   false,
		},
		{
			name:      "success_type_duplicate",
			args:      args{year: 2020, ts: []dayType{TypeWeekDay, TypeWeekDay}},
			wantCount: 246,
			wantErr:   false,
		},
		{
			name:      "too_small_year",
			args:      args{year: minYear - 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: maxYear + 1},
			wantCount: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SpecificTypeDays(tt.args.year, tt.args.ts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpecificTypeDays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Len() != tt.wantCount {
				t.Errorf("ActualCount = %d, wantCount %v", got.Len(), tt.wantCount)
			}
		})
	}
}

func TestSpecificTypeDaysYM(t *testing.T) {
	type args struct {
		year  int
		month int
		ts    []dayType
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "success_weekday",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeWeekDay}},
			wantCount: 21,
			wantErr:   false,
		},
		{
			name:      "success_saturday",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeSaturday}},
			wantCount: 4,
			wantErr:   false,
		},
		{
			name:      "success_sunday",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeSunday}},
			wantCount: 4,
			wantErr:   false,
		},
		{
			name:      "success_holiday",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeNationalHoliday}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:      "success_weekday_and_saturday",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeWeekDay, TypeSaturday}},
			wantCount: 25,
			wantErr:   false,
		},
		{
			name:      "success_sunday_and_holiday",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeSunday, TypeNationalHoliday}},
			wantCount: 6,
			wantErr:   false,
		},
		{
			name:      "success_all",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeWeekDay, TypeSaturday, TypeSunday, TypeNationalHoliday}},
			wantCount: 31,
			wantErr:   false,
		},
		{
			name:      "success_type_duplicate",
			args:      args{year: 2020, month: 1, ts: []dayType{TypeWeekDay, TypeWeekDay}},
			wantCount: 21,
			wantErr:   false,
		},
		{
			name:      "too_small_year",
			args:      args{year: minYear - 1, month: 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: maxYear + 1, month: 12},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_small_month",
			args:      args{year: 2020, month: 0},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_month",
			args:      args{year: 2020, month: 13},
			wantCount: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SpecificTypeDaysYM(tt.args.year, tt.args.month, tt.args.ts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpecificTypeDaysYM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Len() != tt.wantCount {
				t.Errorf("ActualCount = %d, wantCount %v", got.Len(), tt.wantCount)
			}
		})
	}
}
