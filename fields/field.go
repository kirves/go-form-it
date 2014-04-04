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
	label      string
	labelClass []string
	tag        map[string]struct{}
	value      string
	helptext   string
	errors     []string
	multValues map[string]struct{}
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
	SetLabel(label string) FieldInterface
	AddLabelClass(class string) FieldInterface
	RemoveLabelClass(class string) FieldInterface
	SetValue(value string) FieldInterface
	Disabled() FieldInterface
	Enabled() FieldInterface
	SetHelptext(text string) FieldInterface
	AddError(err string) FieldInterface
}

// FieldWithType creates an empty field of the given type and identified by name.
func FieldWithType(name, t string) Field {
	return Field{
		fieldType:  t,
		Widget:     nil,
		name:       name,
		class:      []string{},
		id:         "",
		params:     map[string]string{},
		css:        map[string]string{},
		label:      "",
		labelClass: []string{},
		tag:        map[string]struct{}{},
		value:      "",
		helptext:   "",
		errors:     []string{},
		multValues: map[string]struct{}{},
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

func (f *Field) dataForRender() map[string]interface{} {
	return map[string]interface{}{
		"classes":      f.class,
		"id":           f.id,
		"name":         f.name,
		"params":       f.params,
		"css":          f.css,
		"type":         f.fieldType,
		"label":        f.label,
		"labelClasses": f.labelClass,
		"tags":         f.tag,
		"value":        f.value,
		"helptext":     f.helptext,
		"errors":       f.errors,
	}
}

// Render packs all data and executes widget render method.
func (f *Field) Render() template.HTML {
	if f.Widget != nil {
		data := f.dataForRender()
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
	f.tag[tag] = struct{}{}
	return f
}

// RemoveTag removes a no-value parameter from the field.
func (f *Field) RemoveTag(tag string) FieldInterface {
	delete(f.tag, tag)
	return f
}

// SetValue sets the value parameter for the field.
func (f *Field) SetValue(value string) FieldInterface {
	f.value = value
	return f
}

// SetHelptext saves the field helptext.
func (f *Field) SetHelptext(text string) FieldInterface {
	f.helptext = text
	return f
}

// AddError adds an error string to the field. It's valid only for Bootstrap forms.
func (f *Field) AddError(err string) FieldInterface {
	f.errors = append(f.errors, err)
	return f
}
