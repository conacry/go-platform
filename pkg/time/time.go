package time

import (
	goTime "time"
)

const (
	// Форматирование даты в формате '22.03.2019'
	LayoutDateDMYByPoint = "02.01.2006"
	// Форматирование даты в формате '22/03/2019'
	LayoutDateDMYBySlash = "02/01/2006"
	// Форматирование даты в формате '22/03/2019'
	LayoutDateYMDByDash = "2006-01-02"
	// Форматирование даты и времени в формате '22.03.2019 17:01'
	LayoutDateTimeDMYMMByPoint = "02.01.2006 15:04"
	// Форматирование даты и времени в формате '22.03.2019 17:01:31'
	LayoutDateTimeDMYSSByPoint = "02.01.2006 15:04:05"
)

type Time struct {
	goTime.Time
}

func Now() *Time {
	return &Time{goTime.Now().UTC()}
}

func Empty() *Time {
	return &Time{}
}

func FromTime(t goTime.Time) *Time {
	return &Time{goTime.Unix(0, t.UnixNano()).UTC()}
}

func FromUnixNano(nanoseconds int64) *Time {
	if nanoseconds == 0 {
		return Empty()
	}
	return &Time{goTime.Unix(0, nanoseconds).UTC()}
}

func FromUnixMillis(milliseconds int64) *Time {
	if milliseconds == 0 {
		return Empty()
	}
	return &Time{goTime.UnixMilli(milliseconds).UTC()}
}

func Parse(layout string, value string) (*Time, error) {
	parsedValue, err := goTime.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return FromTime(parsedValue), nil
}

func (t *Time) Local() *Time {
	return &Time{goTime.Unix(0, t.UnixNano()).Local()}
}

func (t *Time) Add(duration goTime.Duration) *Time {
	newTime := t.Time.Add(duration)
	return &Time{newTime}
}

func (t *Time) Sub(duration goTime.Duration) *Time {
	newTime := t.Time.Add(-duration)
	return &Time{newTime}
}

func (t *Time) Equal(time *Time) bool {
	return t.Time.Equal(time.Time)
}

func (t *Time) Before(time *Time) bool {
	return t.Time.Before(time.Time)
}

func (t *Time) After(time *Time) bool {
	return t.Time.After(time.Time)
}

func (t *Time) Unix() int64 {
	var unix int64
	if !t.IsZero() {
		unix = t.Time.Unix()
	}
	return unix
}

func (t *Time) UnixMilli() int64 {
	var unixMilli int64
	if !t.IsZero() {
		unixMilli = t.Time.UnixMilli()
	}
	return unixMilli
}

func (t *Time) UnixNano() int64 {
	var unixNano int64
	if !t.IsZero() {
		unixNano = t.Time.UnixNano()
	}
	return unixNano
}
