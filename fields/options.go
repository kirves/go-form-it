package fields

import (
	"github.com/kirves/revel-forms/common"
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
