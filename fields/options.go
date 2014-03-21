package fields

import (
	"github.com/kirves/revel-forms/widgets"
)

type RadioType struct {
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
