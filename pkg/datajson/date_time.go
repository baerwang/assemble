package datajson

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	FormatDate     = "2006-01-02 15:04:05"
	FormatDateTime = `"2006-01-02T15:04:05.999999"`
)

type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tt.Format(FormatDate))), nil
}

func (t Datetime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).Unix() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t *Datetime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Datetime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *Datetime) UnmarshalJSON(p []byte) error {
	t2, err := time.Parse(FormatDateTime, string(p))
	if err != nil {
		return err
	}
	*t = Datetime(t2)
	return nil
}
