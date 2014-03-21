package fields

import (
	"github.com/kirves/revel-forms/widgets"
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

func TextField(name string) *TextFieldType {
	return &TextFieldType{
		FieldWithType(name, TEXT),
	}
}

func (t *TextFieldType) SetStyle(style string) {
	t.widget = widgets.TextInput(style)
}

func PasswordField(name string) *PasswordFieldType {
	return &PasswordFieldType{
		FieldWithType(name, PASSWORD),
	}
}

func (t *PasswordFieldType) SetStyle(style string) {
	t.widget = widgets.PasswordInput(style)
}

func TextAreaField(name string, rows, cols int) *TextAreaFieldType {
	ret := &TextAreaFieldType{
		FieldWithType(name, TEXTAREA),
	}
	ret.SetParam("rows", string(rows))
	ret.SetParam("cols", string(cols))
	return ret
}

func (t *TextAreaFieldType) SetStyle(style string) {
	t.widget = widgets.TextAreaInput(style)
}
