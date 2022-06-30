package jpcal

import (
	"testing"
)

func TestIsWeekday(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "weekday",
			args:    args{year: 2020, month: 1, day: 2},
			want:    true,
			wantErr: false,
		},
		{
			name:    "saturday",
			args:    args{year: 2020, month: 1, day: 4},
			want:    false,
			wantErr: false,
		},
		{
			name:    "sunday",
			args:    args{year: 2020, month: 1, day: 5},
			want:    false,
			wantErr: false,
		},
		{
			name:    "national holiday",
			args:    args{year: 2020, month: 1, day: 1},
			want:    false,
			wantErr: false,
		},
		{
			name:    "too_small_year",
			args:    args{year: minYear - 1, month: 1, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: maxYear + 1, month: 12, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_small_month",
			args:    args{year: 2020, month: 0, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_month",
			args:    args{year: 2020, month: 13, day: 1},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsWeekday(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsWeekday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSaturday(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "weekday",
			args:    args{year: 2020, month: 1, day: 2},
			want:    false,
			wantErr: false,
		},
		{
			name:    "saturday",
			args:    args{year: 2020, month: 1, day: 4},
			want:    true,
			wantErr: false,
		},
		{
			name:    "sunday",
			args:    args{year: 2020, month: 1, day: 5},
			want:    false,
			wantErr: false,
		},
		{
			name:    "national holiday",
			args:    args{year: 2020, month: 1, day: 1},
			want:    false,
			wantErr: false,
		},
		{
			name:    "too_small_year",
			args:    args{year: minYear - 1, month: 1, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: maxYear + 1, month: 12, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_small_month",
			args:    args{year: 2020, month: 0, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_month",
			args:    args{year: 2020, month: 13, day: 1},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsSaturday(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSaturday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSaturday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSunday(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "weekday",
			args:    args{year: 2020, month: 1, day: 2},
			want:    false,
			wantErr: false,
		},
		{
			name:    "saturday",
			args:    args{year: 2020, month: 1, day: 4},
			want:    false,
			wantErr: false,
		},
		{
			name:    "sunday",
			args:    args{year: 2020, month: 1, day: 5},
			want:    true,
			wantErr: false,
		},
		{
			name:    "national holiday",
			args:    args{year: 2020, month: 1, day: 1},
			want:    false,
			wantErr: false,
		},
		{
			name:    "too_small_year",
			args:    args{year: minYear - 1, month: 1, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: maxYear + 1, month: 12, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_small_month",
			args:    args{year: 2020, month: 0, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_month",
			args:    args{year: 2020, month: 13, day: 1},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsSunday(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSunday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSunday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNationalHoliday(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "weekday",
			args:    args{year: 2020, month: 1, day: 2},
			want:    false,
			wantErr: false,
		},
		{
			name:    "saturday",
			args:    args{year: 2020, month: 1, day: 4},
			want:    false,
			wantErr: false,
		},
		{
			name:    "sunday",
			args:    args{year: 2020, month: 1, day: 5},
			want:    false,
			wantErr: false,
		},
		{
			name:    "national holiday",
			args:    args{year: 2020, month: 1, day: 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "too_small_year",
			args:    args{year: minYear - 1, month: 1, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_year",
			args:    args{year: maxYear + 1, month: 12, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_small_month",
			args:    args{year: 2020, month: 0, day: 1},
			want:    false,
			wantErr: true,
		},
		{
			name:    "too_big_month",
			args:    args{year: 2020, month: 13, day: 1},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsNationalHoliday(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsNationalHoliday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsNationalHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}
