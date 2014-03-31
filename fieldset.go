package forms

import (
	"bytes"
	"fmt"
	"github.com/kirves/revel-forms/fields"
	"html/template"
)

type FieldSetType struct {
	name   string
	class  map[string]struct{}
	tags   map[string]struct{}
	fields []fields.FieldInterface
}

func (f *FieldSetType) Render() template.HTML {
	var s string
	buf := bytes.NewBufferString(s)
	data := map[string]interface{}{
		"fields":  f.fields,
		"classes": f.class,
		"tags":    f.tags,
	}
	err := template.Must(template.ParseFiles("templates/fieldset.html")).Execute(buf, data)
	fmt.Println("NUMBER OF FIELDS:", f.fields)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

func FieldSet(name string, elems ...fields.FieldInterface) *FieldSetType {
	return &FieldSetType{
		name,
		map[string]struct{}{},
		map[string]struct{}{},
		elems,
	}
}

func (f *FieldSetType) Name() string {
	return f.name
}

func (f *FieldSetType) AddClass(class string) *FieldSetType {
	f.class[class] = struct{}{}
	return f
}

func (f *FieldSetType) RemoveClass(class string) *FieldSetType {
	delete(f.class, class)
	return f
}

func (f *FieldSetType) AddTag(tag string) *FieldSetType {
	f.tags[tag] = struct{}{}
	return f
}

func (f *FieldSetType) RemoveTag(tag string) *FieldSetType {
	delete(f.tags, tag)
	return f
}

func (f *FieldSetType) Disable() *FieldSetType {
	f.AddTag("disabled")
	return f
}

func (f *FieldSetType) Enable() *FieldSetType {
	f.RemoveTag("disabled")
	return f
}
