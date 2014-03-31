package fields

import (
	"fmt"
	"github.com/kirves/revel-forms/common"
	"reflect"
	"strconv"
)

type TextFieldType struct {
	Field
}

type PasswordFieldType struct {
	Field
}

type TextAreaFieldType struct {
	Field
}

type HiddenFieldType struct {
	Field
}

func TextField(name string) *TextFieldType {
	return &TextFieldType{
		FieldWithType(name, formcommon.TEXT),
	}
}

func PasswordField(name string) *PasswordFieldType {
	return &PasswordFieldType{
		FieldWithType(name, formcommon.PASSWORD),
	}
}

func TextAreaField(name string, rows, cols int) *TextAreaFieldType {
	ret := &TextAreaFieldType{
		FieldWithType(name, formcommon.TEXTAREA),
	}
	ret.SetParam("rows", string(rows))
	ret.SetParam("cols", string(cols))
	return ret
}

func HiddenField(name string) *HiddenFieldType {
	return &HiddenFieldType{
		FieldWithType(name, formcommon.HIDDEN),
	}
}

func TextFieldFromInstance(i interface{}, fieldNo int, name string) *TextFieldType {
	ret := TextField(name)
	ret.SetValue(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}

func PasswordFieldFromInstance(i interface{}, fieldNo int, name string) *PasswordFieldType {
	ret := PasswordField(name)
	ret.SetValue(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}

func TextAreaFieldFromInstance(i interface{}, fieldNo int, name string) *TextAreaFieldType {
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	var rows, cols int64 = 30, 50
	var err error
	if v := t.Get("form_rows"); v != "" {
		rows, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
	}
	if v := t.Get("form_col"); v != "" {
		cols, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
	}
	ret := TextAreaField(name, int(rows), int(cols))
	ret.SetText(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}

func HiddenFieldFromInstance(i interface{}, fieldNo int, name string) *HiddenFieldType {
	ret := HiddenField(name)
	ret.SetValue(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}
