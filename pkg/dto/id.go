package dto

import (
	"assemble/pkg/utils"
	"database/sql/driver"
)

import (
	"fmt"
	"strconv"
)

type ID int64

func NewID(id int64) ID {
	return ID(id)
}

func NewSnowID() ID {
	return NewID(utils.GoId())
}

func (i *ID) Int64() int64 {
	return int64(*i)
}

func (i *ID) String() string {
	return strconv.FormatInt(i.Int64(), 10)
}

func (i *ID) MarshalJSON() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *ID) UnmarshalJSON(p []byte) error {
	if len(p) <= 2 {
		*i = 0
		return nil
	}

	if p[0] != '"' || p[len(p)-1] != '"' {
		return fmt.Errorf("can't parsing %s to string", string(p))
	}

	str := string(p[1 : len(p)-1])
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*i = ID(i64)
	return nil
}

func (i *ID) Value() (driver.Value, error) {
	return i.Int64(), nil
}

func (i *ID) Scan(v interface{}) error {
	value, ok := v.(int64)
	if ok {
		*i = ID(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to ID", v)
}
