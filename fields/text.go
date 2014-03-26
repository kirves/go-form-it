package fields

import (
	"github.com/kirves/revel-forms/common"
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

// func (t *TextFieldType) SetStyle(style string) {
// 	t.widget = widgets.TextInput(style)
// }

func PasswordField(name string) *PasswordFieldType {
	return &PasswordFieldType{
		FieldWithType(name, formcommon.PASSWORD),
	}
}

// func (t *PasswordFieldType) SetStyle(style string) {
// 	t.widget = widgets.PasswordInput(style)
// }

func TextAreaField(name string, rows, cols int) *TextAreaFieldType {
	ret := &TextAreaFieldType{
		FieldWithType(name, formcommon.TEXTAREA),
	}
	ret.SetParam("rows", string(rows))
	ret.SetParam("cols", string(cols))
	return ret
}

// func (t *TextAreaFieldType) SetStyle(style string) {
// 	t.widget = widgets.TextAreaInput(style)
// }

func HiddenField(name string) *HiddenFieldType {
	return &HiddenFieldType{
		FieldWithType(name, formcommon.HIDDEN),
	}
}

// func (t *HiddenFieldType) SetStyle(style string) {
// 	t.widget = widgets.GenericWidget(style)
// }
