// This package provides form creation and rendering functionalities, as well as FieldSet definition.
// Two kind of forms can be created: base forms and Bootstrap3 compatible forms; even though the latters are automatically provided
// the required classes to make them render correctly in a Bootstrap environment, every form can be given custom parameters such as
// classes, id, generic parameters (in key-value form) and stylesheet options.
package forms

import (
	"github.com/kirves/go-form-it/common"
	"github.com/kirves/go-form-it/fields"
	"html/template"
	"reflect"
	"strings"
)

// Form methods: POST or GET.
const (
	POST = "POST"
	GET  = "GET"
)

// Form structure.
type Form struct {
	fields   []FormElement
	fieldMap map[string]int
	style    string
	template *template.Template
	class    []string
	id       string
	params   map[string]string
	css      map[string]string
	method   string
	action   template.HTML
}

// BaseForm creates an empty form with no styling.
func BaseForm(method, action string) *Form {
	tmpl, err := template.ParseFiles(formcommon.CreateUrl("templates/baseform.html"))
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]FormElement, 0),
		make(map[string]int),
		formcommon.BASE,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		template.HTML(action),
	}
}

// BootstrapForm creates an empty form compliant with Bootstrap3 CSS, both in structure and classes.
func BootstrapForm(method, action string) *Form {
	tmpl, err := template.ParseFiles(formcommon.CreateUrl("templates/baseform.html"))
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]FormElement, 0),
		make(map[string]int),
		formcommon.BOOTSTRAP,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		template.HTML(action),
	}
}

// BaseFormFromModel returns a base form inferring fields, data types and contents from the provided instance.
// A Submit button is automatically added as a last field; the form is editable and fields can be added, modified or removed as needed.
// Tags can be used to drive automatic creation: change default widgets for each field, skip fields or provide additional parameters.
// Basic field -> widget mapping is as follows: string -> textField, bool -> checkbox, time.Time -> datetimeField, int -> numberField;
// nested structs are also converted and added to the form.
func BaseFormFromModel(m interface{}, method, action string) *Form {
	form := BaseForm(method, action)
	for _, v := range unWindStructure(m, "") {
		form.Elements(v)
	}
	form.Elements(fields.SubmitButton("submit", "Submit"))
	return form
}

// Same as BaseFormFromModel but returns a Bootstrap3 compatible form.
func BootstrapFormFromModel(m interface{}, method, action string) *Form {
	form := BootstrapForm(method, action)
	for _, v := range unWindStructure(m, "") {
		form.Elements(v)
	}
	form.Elements(fields.SubmitButton("submit", "Submit"))
	return form
}

func unWindStructure(m interface{}, baseName string) []fields.FieldInterface {
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	fieldList := make([]fields.FieldInterface, 0)
	for i := 0; i < t.NumField(); i++ {
		optionsArr := strings.Split(t.Field(i).Tag.Get("form_options"), ",")
		options := make(map[string]struct{})
		for _, opt := range optionsArr {
			if opt != "" {
				options[opt] = struct{}{}
			}
		}
		if _, ok := options["skip"]; !ok {
			widget := t.Field(i).Tag.Get("form_widget")
			var f fields.FieldInterface
			var fName string
			if baseName == "" {
				fName = t.Field(i).Name
			} else {
				fName = strings.Join([]string{baseName, t.Field(i).Name}, ".")
			}
			switch widget {
			case "text":
				f = fields.TextFieldFromInstance(m, i, fName)
			case "textarea":
				f = fields.TextAreaFieldFromInstance(m, i, fName)
			case "password":
				f = fields.PasswordFieldFromInstance(m, i, fName)
			case "select":
				f = fields.SelectFieldFromInstance(m, i, fName, options)
			case "date":
				f = fields.DateFieldFromInstance(m, i, fName)
			case "datetime":
				f = fields.DatetimeFieldFromInstance(m, i, fName)
			case "time":
				f = fields.TimeFieldFromInstance(m, i, fName)
			case "number":
				f = fields.NumberFieldFromInstance(m, i, fName)
			case "range":
				f = fields.RangeFieldFromInstance(m, i, fName)
			case "radio":
				f = fields.RadioFieldFromInstance(m, i, fName)
			case "static":
				f = fields.StaticFieldFromInstance(m, i, fName)
			default:
				switch t.Field(i).Type.String() {
				case "string":
					f = fields.TextFieldFromInstance(m, i, fName)
				case "bool":
					f = fields.CheckboxFromInstance(m, i, fName, options)
				case "time.Time":
					f = fields.DatetimeFieldFromInstance(m, i, fName)
				case "int":
					f = fields.NumberFieldFromInstance(m, i, fName)
				case "float":
					f = fields.NumberFieldFromInstance(m, i, fName)
				case "struct":
					fieldList = append(fieldList, unWindStructure(v.Field(i).Interface(), fName)...)
					f = nil
				default:
					f = fields.TextFieldFromInstance(m, i, fName)
				}
			}
			if f != nil {
				label := t.Field(i).Tag.Get("form_label")
				if label != "" {
					f.SetLabel(label)
				} else {
					f.SetLabel(strings.Title(t.Field(i).Name))
				}
				fieldList = append(fieldList, f)
			}
		}
	}
	return fieldList
}
