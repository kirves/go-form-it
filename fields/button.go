package fields

import (
	"github.com/kirves/go-form-it/common"
)

// Generic button type
type ButtonType struct {
	Field
}

// SetText saves the provided text as content of the field, usually a TextAreaField.
func (f *ButtonType) SetText(text string) *ButtonType {
	f.additionalData["text"] = text
	return f
}

// SubmitButton creates a default button with the provided name and text.
func SubmitButton(name string, text string) *ButtonType {
	ret := &ButtonType{
		FieldWithType(name, formcommon.SUBMIT),
	}
	ret.SetText(text)
	return ret
}

// ResetButton creates a default reset button with the provided name and text.
func ResetButton(name string, text string) *ButtonType {
	ret := &ButtonType{
		FieldWithType(name, formcommon.RESET),
	}
	ret.SetText(text)
	return ret
}

// Button creates a default generic button
func Button(name string, text string) *ButtonType {
	ret := &ButtonType{
		FieldWithType(name, formcommon.BUTTON),
	}
	ret.SetText(text)
	return ret
}
