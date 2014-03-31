package fields

import (
	"fmt"
	"github.com/kirves/revel-forms/common"
	"reflect"
	"strings"
)

type InputChoice struct {
	Id, Val string
}

type RadioType struct {
	Field
}

type SelectType struct {
	Field
}

type CheckBoxType struct {
	Field
}

func RadioField(name string, choices []InputChoice) *RadioType {
	ret := &RadioType{
		FieldWithType(name, formcommon.RADIO),
	}
	chMap := map[string][]InputChoice{
		"": choices,
	}
	ret.SetChoices(chMap)
	return ret
}

func SelectField(name string, choices map[string][]InputChoice) *SelectType {
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
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	choices := strings.Split(t.Get("form_choices"), "|")
	chArr := make([]InputChoice, 0)
	chMap := make(map[string]string)
	for i := 0; i < len(choices)-1; i += 2 {
		chArr = append(chArr, InputChoice{choices[i], choices[i+1]})
		chMap[choices[i]] = choices[i+1]
	}
	ret := RadioField(name, chArr)

	var v string = t.Get("form_value")
	if v == "" {
		v = fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String())
	}
	if _, ok := chMap[v]; ok {
		ret.SetValue(v)
	}
	return ret
}

func SelectFieldFromInstance(i interface{}, fieldNo int, name string) *SelectType {
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	choices := strings.Split(t.Get("form_choices"), "|")
	chArr := make(map[string][]InputChoice)
	chMap := make(map[string]string)
	for i := 0; i < len(choices)-2; i += 3 {
		if _, ok := chArr[choices[i]]; !ok {
			chArr[choices[i]] = make([]InputChoice, 0)
		}
		chArr[choices[i]] = append(chArr[choices[i]], InputChoice{choices[i+1], choices[i+2]})
		chMap[choices[i+1]] = choices[i+2]
	}
	ret := SelectField(name, chArr)

	var v string = fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String())
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
