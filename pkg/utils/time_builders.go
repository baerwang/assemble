package utils

import (
	"time"
)

const dateTime = "2006-01-02 15:04:05"

type BuilderTime struct {
	Year  int
	Month time.Month
	Day   int
	Hour  int
	Min   int
	Sec   int
	NSec  int
	Local *time.Location
	Now   time.Time
}

func NewBuilderTime() *BuilderTime {
	return &BuilderTime{Now: time.Now(), Local: time.Local}
}

func NewBuilderParseTime(now time.Time) *BuilderTime {
	return &BuilderTime{Now: now, Local: now.Location()}
}

func (b *BuilderTime) SetYear(y int) *BuilderTime {
	b.Year = y
	return b
}

func (b *BuilderTime) SetMonth(m time.Month) *BuilderTime {
	b.Month = m
	return b
}

func (b *BuilderTime) SetDay(d int) *BuilderTime {
	b.Day = d
	return b
}

func (b *BuilderTime) SetHour(h int) *BuilderTime {
	b.Hour = h
	return b
}

func (b *BuilderTime) SetMin(m int) *BuilderTime {
	b.Min = m
	return b
}

func (b *BuilderTime) SetSec(s int) *BuilderTime {
	b.Sec = s
	return b
}

func (b *BuilderTime) SetNSec(n int) *BuilderTime {
	b.NSec = n
	return b
}

func (b *BuilderTime) SetLocal(local *time.Location) *BuilderTime {
	b.Local = local
	return b
}

func (b *BuilderTime) CreateTime() time.Time {
	return time.Date(b.Now.Year()+b.Year, b.Now.Month()+b.Month, b.Now.Day()+b.Day, b.Hour, b.Min, b.Sec, b.NSec, b.Local)
}

func (b *BuilderTime) SubTime(now time.Time) time.Duration {
	return b.CreateTime().Sub(now)
}

func (b *BuilderTime) Format(format string) string {
	return b.CreateTime().Format(format)
}

func (b *BuilderTime) FormatDateTime() string {
	return b.CreateTime().Format(dateTime)
}

// CheckTime 检查时间是不是符合形式
func CheckTime(dateStr string) (res bool) {
	_, err := time.Parse(dateTime, dateStr)
	return err == nil
}

func TimeSupplement(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 999999999, d.Location())
}
