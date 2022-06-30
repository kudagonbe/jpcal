package jpcal

import (
	"reflect"
	"testing"
)

var wantHolidays2020 Days = []Day{
	&nationalHoliday{year: 2020, month: 1, day: 1, holidayName: "元日"},
	&nationalHoliday{year: 2020, month: 1, day: 13, holidayName: "成人の日"},
	&nationalHoliday{year: 2020, month: 2, day: 11, holidayName: "建国記念の日"},
	&nationalHoliday{year: 2020, month: 2, day: 23, holidayName: "天皇誕生日"},
	&nationalHoliday{year: 2020, month: 2, day: 24, holidayName: "休日"},
	&nationalHoliday{year: 2020, month: 3, day: 20, holidayName: "春分の日"},
	&nationalHoliday{year: 2020, month: 4, day: 29, holidayName: "昭和の日"},
	&nationalHoliday{year: 2020, month: 5, day: 3, holidayName: "憲法記念日"},
	&nationalHoliday{year: 2020, month: 5, day: 4, holidayName: "みどりの日"},
	&nationalHoliday{year: 2020, month: 5, day: 5, holidayName: "こどもの日"},
	&nationalHoliday{year: 2020, month: 5, day: 6, holidayName: "休日"},
	&nationalHoliday{year: 2020, month: 7, day: 23, holidayName: "海の日"},
	&nationalHoliday{year: 2020, month: 7, day: 24, holidayName: "スポーツの日"},
	&nationalHoliday{year: 2020, month: 8, day: 10, holidayName: "山の日"},
	&nationalHoliday{year: 2020, month: 9, day: 21, holidayName: "敬老の日"},
	&nationalHoliday{year: 2020, month: 9, day: 22, holidayName: "秋分の日"},
	&nationalHoliday{year: 2020, month: 11, day: 3, holidayName: "文化の日"},
	&nationalHoliday{year: 2020, month: 11, day: 23, holidayName: "勤労感謝の日"},
}

func TestNationalHolidays(t *testing.T) {
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
			args:    args{year: MinYear - 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: MaxYear + 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NationalHolidays(tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("NationalHolidays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NationalHolidays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNationalHolidaysYM(t *testing.T) {
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
			args:    args{year: MinYear - 1, month: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: MaxYear + 1, month: 12},
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
			got, err := NationalHolidaysYM(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("NationalHolidaysYM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NationalHolidaysYM() = %v, want %v", got, tt.want)
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
			args:      args{year: MinYear - 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: MaxYear + 1},
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
			args:      args{year: MinYear - 1, month: 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: MaxYear + 1, month: 12},
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
		ts   []DayType
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "success_weekday",
			args:      args{year: 2020, ts: []DayType{TypeWeekDay}},
			wantCount: 246,
			wantErr:   false,
		},
		{
			name:      "success_saturday",
			args:      args{year: 2020, ts: []DayType{TypeSaturday}},
			wantCount: 52,
			wantErr:   false,
		},
		{
			name:      "success_sunday",
			args:      args{year: 2020, ts: []DayType{TypeSunday}},
			wantCount: 50,
			wantErr:   false,
		},
		{
			name:      "success_holiday",
			args:      args{year: 2020, ts: []DayType{TypeNationalHoliday}},
			wantCount: 18,
			wantErr:   false,
		},
		{
			name:      "success_weekday_and_saturday",
			args:      args{year: 2020, ts: []DayType{TypeWeekDay, TypeSaturday}},
			wantCount: 298,
			wantErr:   false,
		},
		{
			name:      "success_sunday_and_holiday",
			args:      args{year: 2020, ts: []DayType{TypeSunday, TypeNationalHoliday}},
			wantCount: 68,
			wantErr:   false,
		},
		{
			name:      "success_all",
			args:      args{year: 2020, ts: []DayType{TypeWeekDay, TypeSaturday, TypeSunday, TypeNationalHoliday}},
			wantCount: 366,
			wantErr:   false,
		},
		{
			name:      "success_type_duplicate",
			args:      args{year: 2020, ts: []DayType{TypeWeekDay, TypeWeekDay}},
			wantCount: 246,
			wantErr:   false,
		},
		{
			name:      "too_small_year",
			args:      args{year: MinYear - 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: MaxYear + 1},
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
		ts    []DayType
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "success_weekday",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeWeekDay}},
			wantCount: 21,
			wantErr:   false,
		},
		{
			name:      "success_saturday",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeSaturday}},
			wantCount: 4,
			wantErr:   false,
		},
		{
			name:      "success_sunday",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeSunday}},
			wantCount: 4,
			wantErr:   false,
		},
		{
			name:      "success_holiday",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeNationalHoliday}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:      "success_weekday_and_saturday",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeWeekDay, TypeSaturday}},
			wantCount: 25,
			wantErr:   false,
		},
		{
			name:      "success_sunday_and_holiday",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeSunday, TypeNationalHoliday}},
			wantCount: 6,
			wantErr:   false,
		},
		{
			name:      "success_all",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeWeekDay, TypeSaturday, TypeSunday, TypeNationalHoliday}},
			wantCount: 31,
			wantErr:   false,
		},
		{
			name:      "success_type_duplicate",
			args:      args{year: 2020, month: 1, ts: []DayType{TypeWeekDay, TypeWeekDay}},
			wantCount: 21,
			wantErr:   false,
		},
		{
			name:      "too_small_year",
			args:      args{year: MinYear - 1, month: 1},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "too_big_year",
			args:      args{year: MaxYear + 1, month: 12},
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
