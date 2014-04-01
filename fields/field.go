// This package provides all the input fields logic and customization methods.
package fields

import (
	"github.com/kirves/go-form-it/widgets"
	"html/template"
)

// Field is a generic type containing all data associated to an input field.
type Field struct {
	fieldType  string
	Widget     widgets.WidgetInterface // Public Widget field for widget customization
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

// FieldInterface defines the interface an object must implement to be used in a form. Every method returns a FieldInterface object
// to allow methods chaining.
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

// FieldWithType creates an empty field of the given type and identified by name.
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

// SetStyle sets the style (e.g.: BASE, BOOTSTRAP) of the field, correctly populating the Widget field.
func (f *Field) SetStyle(style string) FieldInterface {
	f.Widget = widgets.BaseWidget(style, f.fieldType)
	return f
}

// Name returns the name of the field.
func (f *Field) Name() string {
	return f.name
}

// Render packs all data and executes widget render method.
func (f *Field) Render() template.HTML {
	if f.Widget != nil {
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
		return template.HTML(f.Widget.Render(data))
	}
	return template.HTML("")
}

// AddClass adds a class to the field.
func (f *Field) AddClass(class string) FieldInterface {
	f.class = append(f.class, class)
	return f
}

// RemoveClass removes a class from the field, if it was present.
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

// SetId associates the given id to the field, overwriting any previous id.
func (f *Field) SetId(id string) FieldInterface {
	f.id = id
	return f
}

// SetText saves the provided text as content of the field, usually a TextAreaField.
func (f *Field) SetText(text string) FieldInterface {
	f.text = text
	return f
}

// SetLabel saves the label to be rendered along with the field.
func (f *Field) SetLabel(label string) FieldInterface {
	f.label = label
	return f
}

// SetLablClass allows to define custom classes for the label.
func (f *Field) AddLabelClass(class string) FieldInterface {
	f.labelClass = append(f.labelClass, class)
	return f
}

// RemoveLabelClass removes the given class from the field label.
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

// SetChoices takes as input a dictionary whose key-value entries are defined as follows: key is the group name (the empty string
// is the default group that is not explicitly rendered) and value is the list of choices belonging to that group.
// Grouping is only useful for Select fields, while groups are ignored in Radio fields.
func (f *Field) SetChoices(choices map[string][]InputChoice) FieldInterface {
	f.choices = choices
	return f
}

// SetParam adds a parameter (defined as key-value pair) in the field.
func (f *Field) SetParam(key, value string) FieldInterface {
	f.params[key] = value
	return f
}

// DeleteParam removes a parameter identified by key from the field.
func (f *Field) DeleteParam(key string) FieldInterface {
	delete(f.params, key)
	return f
}

// AddCss adds a custom CSS style the field.
func (f *Field) AddCss(key, value string) FieldInterface {
	f.css[key] = value
	return f
}

// RemoveCss removes CSS options identified by key from the field.
func (f *Field) RemoveCss(key string) FieldInterface {
	delete(f.css, key)
	return f
}

// Disabled add the "disabled" tag to the field, making it unresponsive in some environments (e.g. Bootstrap).
func (f *Field) Disabled() FieldInterface {
	f.AddTag("disabled")
	return f
}

// Enabled removes the "disabled" tag from the field, making it responsive.
func (f *Field) Enabled() FieldInterface {
	f.RemoveTag("disabled")
	return f
}

// AddTag adds a no-value parameter (e.g.: checked, disabled) to the field.
func (f *Field) AddTag(tag string) FieldInterface {
	f.tag = append(f.tag, tag)
	return f
}

// RemoveTag removes a no-value parameter from the field.
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

// SetValue sets the value parameter for the field.
func (f *Field) SetValue(value string) FieldInterface {
	f.value = value
	return f
}
