package fields

import (
	"github.com/kirves/go-form-it/widgets"
	"html/template"
)

type Field struct {
	fieldType  string
	widget     widgets.WidgetInterface
	name       string
	class      []string
	id         string
	params     map[string]string
	css        map[string]string
	text       string
	label      string
	choices    map[string][]InputChoice
	labelClass []string
	tag        []string
	value      string
}

type FieldInterface interface {
	Name() string
	Render() template.HTML
	AddClass(class string) FieldInterface
	RemoveClass(class string) FieldInterface
	AddTag(class string) FieldInterface
	RemoveTag(class string) FieldInterface
	SetId(id string) FieldInterface
	SetParam(key, value string) FieldInterface
	DeleteParam(key string) FieldInterface
	AddCss(key, value string) FieldInterface
	RemoveCss(key string) FieldInterface
	SetStyle(style string) FieldInterface
	SetText(text string) FieldInterface
	SetLabel(label string) FieldInterface
	AddLabelClass(class string) FieldInterface
	RemoveLabelClass(class string) FieldInterface
	SetChoices(choices map[string][]InputChoice) FieldInterface
	SetValue(value string) FieldInterface
	Disabled() FieldInterface
	Enabled() FieldInterface
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
		map[string][]InputChoice{},
		[]string{},
		[]string{},
		"",
	}
}

func (f *Field) SetStyle(style string) FieldInterface {
	f.widget = widgets.BaseWidget(style, f.fieldType)
	return f
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Render() template.HTML {
	if f.widget != nil {
		data := map[string]interface{}{
			"classes":      f.class,
			"id":           f.id,
			"name":         f.name,
			"params":       f.params,
			"css":          f.css,
			"text":         f.text,
			"type":         f.fieldType,
			"label":        f.label,
			"choices":      f.choices,
			"labelClasses": f.labelClass,
			"tags":         f.tag,
			"value":        f.value,
		}
		return template.HTML(f.widget.Render(data))
	}
	return template.HTML("")
}

func (f *Field) AddClass(class string) FieldInterface {
	f.class = append(f.class, class)
	return f
}

func (f *Field) RemoveClass(class string) FieldInterface {
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
	return f
}

func (f *Field) SetId(id string) FieldInterface {
	f.id = id
	return f
}

func (f *Field) SetText(text string) FieldInterface {
	f.text = text
	return f
}

func (f *Field) SetLabel(label string) FieldInterface {
	f.label = label
	return f
}

func (f *Field) AddLabelClass(class string) FieldInterface {
	f.labelClass = append(f.labelClass, class)
	return f
}

func (f *Field) RemoveLabelClass(class string) FieldInterface {
	ind := -1
	for i, v := range f.labelClass {
		if v == class {
			ind = i
			break
		}
	}

	if ind != -1 {
		f.labelClass = append(f.labelClass[:ind], f.labelClass[ind+1:]...)
	}
	return f
}

func (f *Field) SetChoices(choices map[string][]InputChoice) FieldInterface {
	f.choices = choices
	return f
}

func (f *Field) SetParam(key, value string) FieldInterface {
	f.params[key] = value
	return f
}

func (f *Field) DeleteParam(key string) FieldInterface {
	delete(f.params, key)
	return f
}

func (f *Field) AddCss(key, value string) FieldInterface {
	f.css[key] = value
	return f
}

func (f *Field) RemoveCss(key string) FieldInterface {
	delete(f.css, key)
	return f
}

func (f *Field) Disabled() FieldInterface {
	f.AddTag("disabled")
	return f
}

func (f *Field) Enabled() FieldInterface {
	f.RemoveTag("disabled")
	return f
}

func (f *Field) AddTag(tag string) FieldInterface {
	f.tag = append(f.tag, tag)
	return f
}

func (f *Field) RemoveTag(tag string) FieldInterface {
	ind := -1
	for i, v := range f.tag {
		if v == tag {
			ind = i
			break
		}
	}

	if ind != -1 {
		f.tag = append(f.tag[:ind], f.tag[ind+1:]...)
	}
	return f
}

func (f *Field) SetValue(value string) FieldInterface {
	f.value = value
	return f
}
