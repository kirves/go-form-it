package fields

import (
	"github.com/kirves/go-form-it/common"
	"html/template"
)

// Generic button type
type ButtonType struct {
	Field
	text string
}

// SetText saves the provided text as content of the field, usually a TextAreaField.
func (f *ButtonType) SetText(text string) *ButtonType {
	f.text = text
	return f
}

// Render packs all data and executes widget render method.
func (f *ButtonType) Render() template.HTML {
	if f.Widget != nil {
		data := f.dataForRender()
		data["text"] = f.text
		return template.HTML(f.Widget.Render(data))
	}
	return template.HTML("")
}

// // Generic button type
// type ButtonType tButton

// // Submit button type
// type SubmitButtonType tButton

// // Reset button type
// type ResetButtonType tButton

// SubmitButton creates a default button with the provided name and text.
func SubmitButton(name string, text string) *ButtonType {
	ret := &ButtonType{
		FieldWithType(name, formcommon.SUBMIT),
		"",
	}
	ret.SetText(text)
	return ret
}

// ResetButton creates a default reset button with the provided name and text.
func ResetButton(name string, text string) *ButtonType {
	ret := &ButtonType{
		FieldWithType(name, formcommon.RESET),
		"",
	}
	ret.SetText(text)
	return ret
}

// Button creates a default generic button
func Button(name string, text string) *ButtonType {
	ret := &ButtonType{
		FieldWithType(name, formcommon.BUTTON),
		"",
	}
	ret.SetText(text)
	return ret
}
