package forms

import (
	"github.com/kirves/revel-forms/common"
	"github.com/kirves/revel-forms/fields"
	"html/template"
	"reflect"
	"strings"
)

const (
	POST = "POST"
	GET  = "GET"
)

type Form struct {
	fields   []fields.FieldInterface
	fieldMap map[string]int
	style    string
	template *template.Template
	class    []string
	id       string
	params   map[string]string
	css      map[string]string
	method   string
	action   string
}

func BaseForm(method, action string) *Form {
	tmpl, err := template.ParseFiles("templates/baseform.html")
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]fields.FieldInterface, 0),
		make(map[string]int),
		formcommon.BASE,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		action,
	}
}

func BootstrapForm(method, action string) *Form {
	tmpl, err := template.ParseFiles("templates/bootstrapform.html")
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]fields.FieldInterface, 0),
		make(map[string]int),
		formcommon.BOOTSTRAP,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		action,
	}
}

func BaseFormFromModel(m interface{}, method, action string) *Form {
	form := BaseForm(method, action)
	for _, v := range unWindStructure(m, "") {
		form.AddField(v)
	}
	form.AddField(fields.SubmitButton("submit", "Submit"))
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
				f = fields.SelectFieldFromInstance(m, i, fName)
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
