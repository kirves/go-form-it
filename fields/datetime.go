package fields

import (
	"errors"
	"fmt"
	"github.com/kirves/go-form-it/common"
	"reflect"
	"time"
)

// Datetime format string to convert from time.Time objects to HTML fields and viceversa.
const (
	DATETIME_FORMAT = "2006-01-02T15:05"
	DATE_FORMAT     = "2006-01-02"
	TIME_FORMAT     = "15:05"
)

// // DateTime type.
// type DatetimeType struct {
// 	Field
// }

// // Date type.
// type DateType struct {
// 	Field
// }

// // Time type.
// type TimeType struct {
// 	Field
// }

// DatetimeField creates a default datetime input field with the given name.
func DatetimeField(name string) *Field {
	ret := FieldWithType(name, formcommon.DATETIME)
	return ret
}

// DateField creates a default date input field with the given name.
func DateField(name string) *Field {
	ret := FieldWithType(name, formcommon.DATE)
	return ret
}

// TimeField creates a default time input field with the given name.
func TimeField(name string) *Field {
	ret := FieldWithType(name, formcommon.TIME)
	return ret
}

// DatetimeFieldFromInstance creates and initializes a datetime field based on its name, the reference object instance and field number.
// This method looks for "form_min", "form_max" and "form_value" tags to add additional parameters to the field.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func DatetimeFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := DatetimeField(name)
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

// DateFieldFromInstance creates and initializes a date field based on its name, the reference object instance and field number.
// This method looks for "form_min", "form_max" and "form_value" tags to add additional parameters to the field.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func DateFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := DateField(name)
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

// TimeFieldFromInstance creates and initializes a time field based on its name, the reference object instance and field number.
// This method looks for "form_min", "form_max" and "form_value" tags to add additional parameters to the field.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func TimeFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := TimeField(name)
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
