package fields

import (
	"errors"
	"fmt"
	"github.com/kirves/revel-forms/common"
	"reflect"
	"time"
)

const (
	DATETIME_FORMAT = "2006-01-02T15:05"
	DATE_FORMAT     = "2006-01-02"
	TIME_FORMAT     = "15:05"
)

type DatetimeType struct {
	Field
}

type DateType struct {
	Field
}

type TimeType struct {
	Field
}

func DatetimeField(name string) *DatetimeType {
	ret := &DatetimeType{
		FieldWithType(name, formcommon.DATETIME),
	}
	return ret
}

func DateField(name string) *DateType {
	ret := &DateType{
		FieldWithType(name, formcommon.DATE),
	}
	return ret
}

func TimeField(name string) *TimeType {
	ret := &TimeType{
		FieldWithType(name, formcommon.TIME),
	}
	return ret
}

func DatetimeFieldFromInstance(i interface{}, fieldNo int, name string) *DatetimeType {
	ret := &DatetimeType{
		FieldWithType(name, formcommon.DATETIME),
	}
	// check tags
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	if v := t.Get("form_min"); v != "" {
		if !validateDatetime(v) {
			panic(errors.New(fmt.Sprintf("Invalid date value (min) for field: %s", name)))
		}
		ret.SetParam("min", v)
	}
	if v := t.Get("form_max"); v != "" {
		if !validateDatetime(v) {
			panic(errors.New(fmt.Sprintf("Invalid date value (max) for field: %s", name)))
		}
		ret.SetParam("max", v)
	}
	if v := t.Get("form_value"); v != "" {
		ret.SetValue(v)
	} else {
		if v := reflect.ValueOf(i).Field(fieldNo).Interface().(time.Time); !v.IsZero() {
			ret.SetValue(v.Format(DATETIME_FORMAT))
		}
	}
	return ret
}

func DateFieldFromInstance(i interface{}, fieldNo int, name string) *DateType {
	ret := &DateType{
		FieldWithType(name, formcommon.DATE),
	}
	// check tags
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	if v := t.Get("form_min"); v != "" {
		if !validateDate(v) {
			panic(errors.New(fmt.Sprintf("Invalid date value (min) for field", name)))
		}
		ret.SetParam("min", v)
	}
	if v := t.Get("form_max"); v != "" {
		if !validateDate(v) {
			panic(errors.New(fmt.Sprintf("Invalid date value (max) for field", name)))
		}
		ret.SetParam("max", v)
	}
	if v := t.Get("form_value"); v != "" {
		ret.SetValue(v)
	} else {
		if v := reflect.ValueOf(i).Field(fieldNo).Interface().(time.Time); !v.IsZero() {
			ret.SetValue(v.Format(DATE_FORMAT))
		}
	}
	return ret
}

func TimeFieldFromInstance(i interface{}, fieldNo int, name string) *TimeType {
	ret := &TimeType{
		FieldWithType(name, formcommon.TIME),
	}
	// check tags
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	if v := t.Get("form_min"); v != "" {
		if !validateTime(v) {
			panic(errors.New(fmt.Sprintf("Invalid time value (min) for field", name)))
		}
		ret.SetParam("min", v)
	}
	if v := t.Get("form_max"); v != "" {
		if !validateTime(v) {
			panic(errors.New(fmt.Sprintf("Invalid time value (max) for field", name)))
		}
		ret.SetParam("max", v)
	}
	if v := t.Get("form_value"); v != "" {
		ret.SetValue(v)
	} else {
		if v := reflect.ValueOf(i).Field(fieldNo).Interface().(time.Time); !v.IsZero() {
			ret.SetValue(v.Format(TIME_FORMAT))
		}
	}
	return ret
}

func validateDatetime(v string) bool {
	_, err := time.Parse(DATETIME_FORMAT, v)
	if err != nil {
		return false
	}
	return true
}

func validateDate(v string) bool {
	_, err := time.Parse(DATE_FORMAT, v)
	if err != nil {
		return false
	}
	return true
}

func validateTime(v string) bool {
	_, err := time.Parse(TIME_FORMAT, v)
	if err != nil {
		return false
	}
	return true
}
