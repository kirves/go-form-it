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
		tag := t.Field(i).Tag.Get("form_skip")
		if tag == "" {
			tag = t.Field(i).Tag.Get("form_widget")
			var f fields.FieldInterface
			var fName string
			if baseName == "" {
				fName = t.Field(i).Name
			} else {
				fName = strings.Join([]string{baseName, t.Field(i).Name}, ".")
			}
			switch tag {
			case "text":
				f = fields.TextField(fName)
			case "textarea":
				f = fields.TextAreaField(fName, 30, 50)
			case "password":
				f = fields.PasswordField(fName)
			case "select":
				choices := strings.Split(t.Field(i).Tag.Get("form_choices"), "|")
				if len(choices)%2 != 0 {
					f = nil
					break
				}
				chMap := make(map[string]string)
				for i := 0; i < len(choices)-1; i += 2 {
					chMap[choices[i]] = choices[i+1]
				}
				f = fields.SelectField(fName, chMap)
			case "date":
				f = fields.DateField(fName)
			case "datetime":
				f = fields.DatetimeFieldFromInstance(m, i, fName)
			case "time":
				f = fields.TimeField(fName)
			case "number":
				f = fields.NumberFieldFromInstance(m, i, fName)
			case "range":
			default:
				switch t.Field(i).Type.String() {
				case "string":
					f = fields.TextField(fName)
				case "bool":
					initVal := t.Field(i).Tag.Get("form_checked")
					if initVal != "" {
						f = fields.Checkbox(fName, true)
					} else {
						f = fields.Checkbox(fName, false)
					}
				case "time.Time":
					f = fields.DatetimeFieldFromInstance(m, i, fName)
				case "int":
					f = fields.NumberFieldFromInstance(m, i, fName)
				case "struct":
					fieldList = append(fieldList, unWindStructure(v.Field(i).Interface(), fName)...)
					f = nil
				default:
					f = fields.TextField(fName)
				}
			}
			if f != nil {
				f.SetLabel(strings.Title(t.Field(i).Name))
				fieldList = append(fieldList, f)
			}
		}
	}
	return fieldList
}
