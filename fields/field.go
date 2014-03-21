package fields

import (
	"github.com/kirves/revel-forms/widgets"
	"html/template"
)

const (
	BUTTON         = "button" //
	CHECKBOX       = "checkbox"
	COLOR          = "color"
	DATE           = "date"
	DATETIME       = "datetime"
	DATETIME_LOCAL = "datetime-local"
	EMAIL          = "email"
	FILE           = "file"
	HIDDEN         = "hidden"
	IMAGE          = "image"
	MONTH          = "month"
	NUMBER         = "number"
	PASSWORD       = "password" //
	RADIO          = "radio"
	RANGE          = "range"
	RESET          = "reset" //
	SEARCH         = "search"
	SUBMIT         = "submit" //
	TEL            = "tel"
	TEXT           = "text" //
	TIME           = "time"
	URL            = "url"
	WEEK           = "week"
	TEXTAREA       = "textarea" //
	SELECT         = "select"
)

const (
	BASE      = "base"
	BOOTSTRAP = "bootstrap3"
)

type Field struct {
	fieldType string
	widget    widgets.WidgetInterface
	name      string
	class     []string
	id        string
	params    map[string]string
	css       map[string]string
	text      string
	label     string
	choices   map[string]string
}

type FieldInterface interface {
	Name() string
	Render() template.HTML
	AddClass(class string)
	RemoveClass(class string)
	SetId(id string)
	SetParam(key, value string)
	DeleteParam(key string)
	AddCss(key, value string)
	RemoveCss(key string)
	SetStyle(style string)
	SetText(text string)
	SetLabel(label string)
	SetChoices(choices map[string]string)
}

func FieldWithType(name, t string) Field {
	return Field{
		t,
		nil,
		name,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		"",
		"",
		map[string]string{},
	}
}

func (f *Field) SetStyle(style string) {
	return
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Render() template.HTML {
	if f.widget != nil {
		data := map[string]interface{}{
			"classes": f.class,
			"id":      f.id,
			"name":    f.name,
			"params":  f.params,
			"css":     f.css,
			"text":    f.text,
			"type":    f.fieldType,
			"label":   f.label,
			"choices": f.choices,
		}
		return template.HTML(f.widget.Render(data))
	}
	return template.HTML("")
}

func (f *Field) AddClass(class string) {
	f.class = append(f.class, class)
}

func (f *Field) RemoveClass(class string) {
	ind := -1
	for i, v := range f.class {
		if v == class {
			ind = i
			break
		}
	}

	if ind != -1 {
		f.class = append(f.class[:ind], f.class[ind+1:]...)
	}
}

func (f *Field) SetId(id string) {
	f.id = id
}

func (f *Field) SetText(text string) {
	f.text = text
}

func (f *Field) SetLabel(label string) {
	f.label = label
}

func (f *Field) SetChoices(choices map[string]string) {
	f.choices = choices
}

func (f *Field) SetParam(key, value string) {
	f.params[key] = value
}

func (f *Field) DeleteParam(key string) {
	delete(f.params, key)
}

func (f *Field) AddCss(key, value string) {
	f.css[key] = value
}

func (f *Field) RemoveCss(key string) {
	delete(f.css, key)
}
