package forms

import (
	"bytes"
	"fmt"
	"github.com/kirves/go-form-it/common"
	"github.com/kirves/go-form-it/fields"
	"html/template"
)

// FieldSetType is a collection of fields grouped within a form.
type FieldSetType struct {
	name     string
	class    map[string]struct{}
	tags     map[string]struct{}
	fields   []fields.FieldInterface
	fieldMap map[string]int
}

// Render translates a FieldSetType into HTML code and returns it as a template.HTML object.
func (f *FieldSetType) Render() template.HTML {
	var s string
	buf := bytes.NewBufferString(s)
	data := map[string]interface{}{
		"fields":  f.fields,
		"classes": f.class,
		"tags":    f.tags,
	}
	err := template.Must(template.ParseFiles(formcommon.CreateUrl("templates/fieldset.html"))).Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

// FieldSet creates and returns a new FieldSetType with the given name and list of fields.
// Every method for FieldSetType objects returns the object itself, so that call can be chained.
func FieldSet(name string, elems ...fields.FieldInterface) *FieldSetType {
	ret := &FieldSetType{
		name,
		map[string]struct{}{},
		map[string]struct{}{},
		elems,
		map[string]int{},
	}
	for i, elem := range elems {
		ret.fieldMap[elem.Name()] = i
	}
	return ret
}

// Field returns the field identified by name. It returns an empty field if it is missing.
func (f *FieldSetType) Field(name string) fields.FieldInterface {
	ind, ok := f.fieldMap[name]
	if !ok {
		return &fields.Field{}
	}
	return f.fields[ind].(fields.FieldInterface)
}

// Name returns the name of the fieldset.
func (f *FieldSetType) Name() string {
	return f.name
}

// AddClass saves the provided class for the fieldset.
func (f *FieldSetType) AddClass(class string) *FieldSetType {
	f.class[class] = struct{}{}
	return f
}

// RemoveClass removes the provided class from the fieldset, if it was present. Nothing is done if it was not originally present.
func (f *FieldSetType) RemoveClass(class string) *FieldSetType {
	delete(f.class, class)
	return f
}

// AddTag adds a no-value parameter (e.g.: "disabled", "checked") to the fieldset.
func (f *FieldSetType) AddTag(tag string) *FieldSetType {
	f.tags[tag] = struct{}{}
	return f
}

// RemoveTag removes a tag from the fieldset, if it was present.
func (f *FieldSetType) RemoveTag(tag string) *FieldSetType {
	delete(f.tags, tag)
	return f
}

// Disable adds tag "disabled" to the fieldset, making it unresponsive in some environment (e.g.: Bootstrap).
func (f *FieldSetType) Disable() *FieldSetType {
	f.AddTag("disabled")
	return f
}

// Enable removes tag "disabled" from the fieldset, making it responsive.
func (f *FieldSetType) Enable() *FieldSetType {
	f.RemoveTag("disabled")
	return f
}
