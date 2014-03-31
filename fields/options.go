package fields

import (
	"fmt"
	"github.com/kirves/revel-forms/common"
	"reflect"
	"strings"
)

type RadioType struct {
	Field
}

type SelectType struct {
	Field
}

type CheckBoxType struct {
	Field
}

func RadioField(name string, choices map[string]string) *RadioType {
	ret := &RadioType{
		FieldWithType(name, formcommon.RADIO),
	}
	ret.SetChoices(choices)
	return ret
}

func SelectField(name string, choices map[string]string) *SelectType {
	ret := &SelectType{
		FieldWithType(name, formcommon.SELECT),
	}
	ret.SetChoices(choices)
	return ret
}

func Checkbox(name string, checked bool) *CheckBoxType {
	ret := &CheckBoxType{
		FieldWithType(name, formcommon.CHECKBOX),
	}
	if checked {
		ret.AddTag("checked")
	}
	return ret
}

func RadioFieldFromInstance(i interface{}, fieldNo int, name string) *RadioType {
	ret := &RadioType{
		FieldWithType(name, formcommon.RADIO),
	}
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	choices := strings.Split(t.Get("form_choices"), "|")
	chMap := make(map[string]string)
	for i := 0; i < len(choices)-1; i += 2 {
		chMap[choices[i]] = choices[i+1]
	}
	ret.SetChoices(chMap)

	var v string = t.Get("form_value")
	if v == "" {
		v = fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).Interface())
	}
	if _, ok := chMap[v]; ok {
		ret.SetValue(v)
	}
	return ret
}

func SelectFieldFromInstance(i interface{}, fieldNo int, name string) *SelectType {
	ret := &SelectType{
		FieldWithType(name, formcommon.SELECT),
	}
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	choices := strings.Split(t.Get("form_choices"), "|")
	chMap := make(map[string]string)
	for i := 0; i < len(choices)-1; i += 2 {
		chMap[choices[i]] = choices[i+1]
	}
	ret.SetChoices(chMap)

	var v string = fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).Interface())
	if v == "" {
		v = t.Get("form_value")
	}
	if _, ok := chMap[v]; ok {
		ret.SetValue(v)
	}
	return ret
}

func CheckboxFromInstance(i interface{}, fieldNo int, name string, options map[string]struct{}) *CheckBoxType {
	ret := &CheckBoxType{
		FieldWithType(name, formcommon.CHECKBOX),
	}

	if _, ok := options["checked"]; ok {
		ret.AddTag("checked")
	} else {
		val := reflect.ValueOf(i).Field(fieldNo).Bool()
		if val {
			ret.AddTag("checked")
		}
	}
	return ret
}
