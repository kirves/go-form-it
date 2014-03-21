package fields

import (
	"github.com/kirves/revel-forms/widgets"
)

type RadioType struct {
	Field
}

type SelectType struct {
	Field
}

func RadioField(name string, choices map[string]string) *RadioType {
	ret := &RadioType{
		FieldWithType(name, RADIO),
	}
	ret.SetChoices(choices)
	return ret
}

func (t *RadioType) SetStyle(style string) {
	t.widget = widgets.RadioButton(style)
}

func SelectField(name string, choices map[string]string) *SelectType {
	ret := &SelectType{
		FieldWithType(name, SELECT),
	}
	ret.SetChoices(choices)
	return ret
}

func (t *SelectType) SetStyle(style string) {
	t.widget = widgets.SelectMenu(style)
}
