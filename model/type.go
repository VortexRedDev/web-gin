package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type MyTime time.Time

var patten = "2006-01-02 15:04:05"

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(patten))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format(patten), nil
}
func (t MyTime) ToTime() time.Time {
    return time.Time(t)
}

//func (t *MyTime) Scan(v interface{}) error {
//	if value, ok := v.(time.Time); ok {
//		*t = MyTime(value)
//		return nil
//	}
//	return fmt.Errorf("can not convert %v to timestamp", v)
//}
