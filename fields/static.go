package fields

import (
	"fmt"
	"github.com/kirves/go-form-it/common"
	"html/template"
	"reflect"
)

// Static field type
type StaticType struct {
	Field
	text string
}

// SetText saves the provided text as content of the field, usually a TextAreaField.
func (f *StaticType) SetText(text string) *StaticType {
	f.text = text
	return f
}

// Render packs all data and executes widget render method.
func (f *StaticType) Render() template.HTML {
	if f.Widget != nil {
		data := f.dataForRender()
		data["text"] = f.text
		return template.HTML(f.Widget.Render(data))
	}
	return template.HTML("")
}

// StaticField returns a static field with the provided name and content
func StaticField(name, content string) *StaticType {
	ret := &StaticType{
		FieldWithType(name, formcommon.STATIC),
		"",
	}
	ret.SetText(content)
	return ret
}

// RadioFieldFromInstance creates and initializes a radio field based on its name, the reference object instance and field number.
// It uses i object's [fieldNo]-th field content (if any) to set the field content.
func StaticFieldFromInstance(i interface{}, fieldNo int, name string) *StaticType {
	ret := StaticField(name, fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).Interface()))
	return ret
}
