package jpcal

import (
	"reflect"
	"testing"
	"time"
)

func Test_normalDay_Str(t *testing.T) {
	type fields struct {
		year    int
		month   int
		day     int
		dayType DayType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				year:    2020,
				month:   1,
				day:     2,
				dayType: TypeWeekDay,
			},
			want: "2020-01-02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &normalDay{
				year:    tt.fields.year,
				month:   tt.fields.month,
				day:     tt.fields.day,
				dayType: tt.fields.dayType,
			}
			if got := d.Str(); got != tt.want {
				t.Errorf("normalDay.Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalDay_Time(t *testing.T) {
	type fields struct {
		year    int
		month   int
		day     int
		dayType DayType
	}
	tests := []struct {
		name    string
		fields  fields
		want    time.Time
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				year:    2020,
				month:   1,
				day:     2,
				dayType: TypeWeekDay,
			},
			want:    time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &normalDay{
				year:    tt.fields.year,
				month:   tt.fields.month,
				day:     tt.fields.day,
				dayType: tt.fields.dayType,
			}
			got, err := d.Time()
			if (err != nil) != tt.wantErr {
				t.Errorf("normalDay.Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalDay.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalDay_Type(t *testing.T) {
	type fields struct {
		year    int
		month   int
		day     int
		dayType DayType
	}
	tests := []struct {
		name   string
		fields fields
		want   DayType
	}{
		{
			name: "success_weekday",
			fields: fields{
				year:    2020,
				month:   1,
				day:     2,
				dayType: TypeWeekDay,
			},
			want: TypeWeekDay,
		},
		{
			name: "success_saturday",
			fields: fields{
				year:    2020,
				month:   1,
				day:     4,
				dayType: TypeSaturday,
			},
			want: TypeSaturday,
		},
		{
			name: "success_sunday",
			fields: fields{
				year:    2020,
				month:   1,
				day:     5,
				dayType: TypeSunday,
			},
			want: TypeSunday,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &normalDay{
				year:    tt.fields.year,
				month:   tt.fields.month,
				day:     tt.fields.day,
				dayType: tt.fields.dayType,
			}
			if got := d.Type(); got != tt.want {
				t.Errorf("normalDay.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalDay_Description(t *testing.T) {
	type fields struct {
		year    int
		month   int
		day     int
		dayType DayType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success_weekday",
			fields: fields{
				year:    2020,
				month:   1,
				day:     2,
				dayType: TypeWeekDay,
			},
			want: "平日",
		},
		{
			name: "success_saturday",
			fields: fields{
				year:    2020,
				month:   1,
				day:     4,
				dayType: TypeSaturday,
			},
			want: "土曜日",
		},
		{
			name: "success_sunday",
			fields: fields{
				year:    2020,
				month:   1,
				day:     5,
				dayType: TypeSunday,
			},
			want: "日曜日",
		},
		{
			name: "default",
			fields: fields{
				year:    2020,
				month:   1,
				day:     1,
				dayType: TypeNationalHoliday,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &normalDay{
				year:    tt.fields.year,
				month:   tt.fields.month,
				day:     tt.fields.day,
				dayType: tt.fields.dayType,
			}
			if got := d.Description(); got != tt.want {
				t.Errorf("normalDay.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nationalHoliday_Str(t *testing.T) {
	type fields struct {
		year        int
		month       int
		day         int
		holidayName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				year:        2020,
				month:       1,
				day:         1,
				holidayName: "元日",
			},
			want: "2020-01-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nationalHoliday{
				year:        tt.fields.year,
				month:       tt.fields.month,
				day:         tt.fields.day,
				holidayName: tt.fields.holidayName,
			}
			if got := d.Str(); got != tt.want {
				t.Errorf("nationalHoliday.Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nationalHoliday_Time(t *testing.T) {
	type fields struct {
		year        int
		month       int
		day         int
		holidayName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    time.Time
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				year:        2020,
				month:       1,
				day:         1,
				holidayName: "元日",
			},
			want:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nationalHoliday{
				year:        tt.fields.year,
				month:       tt.fields.month,
				day:         tt.fields.day,
				holidayName: tt.fields.holidayName,
			}
			got, err := d.Time()
			if (err != nil) != tt.wantErr {
				t.Errorf("nationalHoliday.Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nationalHoliday.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nationalHoliday_Type(t *testing.T) {
	type fields struct {
		year        int
		month       int
		day         int
		holidayName string
	}
	tests := []struct {
		name   string
		fields fields
		want   DayType
	}{
		{
			name: "success",
			fields: fields{
				year:        2020,
				month:       1,
				day:         1,
				holidayName: "元日",
			},
			want: TypeNationalHoliday,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nationalHoliday{
				year:        tt.fields.year,
				month:       tt.fields.month,
				day:         tt.fields.day,
				holidayName: tt.fields.holidayName,
			}
			if got := d.Type(); got != tt.want {
				t.Errorf("nationalHoliday.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nationalHoliday_Description(t *testing.T) {
	type fields struct {
		year        int
		month       int
		day         int
		holidayName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				year:        2020,
				month:       1,
				day:         1,
				holidayName: "元日",
			},
			want: "元日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nationalHoliday{
				year:        tt.fields.year,
				month:       tt.fields.month,
				day:         tt.fields.day,
				holidayName: tt.fields.holidayName,
			}
			if got := d.Description(); got != tt.want {
				t.Errorf("nationalHoliday.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDay(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    Day
		wantErr bool
	}{
		{
			name: "success_normalday",
			args: args{
				year:  2020,
				month: 1,
				day:   2,
			},
			want:    &normalDay{year: 2020, month: 1, day: 2, dayType: TypeWeekDay},
			wantErr: false,
		},
		{
			name: "success_national_holiday",
			args: args{
				year:  2020,
				month: 1,
				day:   1,
			},
			want:    &nationalHoliday{year: 2020, month: 1, day: 1, holidayName: "元日"},
			wantErr: false,
		},
		{
			name: "too_small_year",
			args: args{
				year:  minYear - 1,
				month: 1,
				day:   1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "too_big_year",
			args: args{
				year:  maxYear + 1,
				month: 1,
				day:   1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid_argument",
			args: args{
				year:  2020,
				month: 1,
				day:   32,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDay(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
