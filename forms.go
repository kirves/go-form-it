package forms

import (
	"bytes"
	"github.com/kirves/revel-forms/fields"
	"html/template"
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

func (f *Form) Render() template.HTML {
	var s string
	buf := bytes.NewBufferString(s)
	// for _, v := range f.fields {
	// 	buf.WriteString(v.Render())
	// 	buf.WriteRune('\n')
	// }
	data := map[string]interface{}{
		"fields":  f.fields,
		"classes": f.class,
		"id":      f.id,
		"params":  f.params,
		"css":     f.css,
		"method":  f.method,
		"action":  f.action,
	}
	err := f.template.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

func BaseForm(method, action string) *Form {
	tmpl, err := template.ParseFiles("templates/baseform.html")
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]fields.FieldInterface, 0),
		make(map[string]int),
		fields.BASE,
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
		fields.BOOTSTRAP,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		action,
	}
}

func (f *Form) AddField(field fields.FieldInterface) fields.FieldInterface {
	field.SetStyle(f.style)
	f.fields = append(f.fields, field)
	f.fieldMap[field.Name()] = len(f.fields) - 1
	return field
}

func (f *Form) RemoveField(name string) {
	ind, ok := f.fieldMap[name]
	if !ok {
		return
	}
	delete(f.fieldMap, name)
	f.fields = append(f.fields[:ind], f.fields[ind+1:]...)
}

func (f *Form) AddClass(class string) *Form {
	f.class = append(f.class, class)
	return f
}

func (f *Form) RemoveClass(class string) *Form {
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

func (f *Form) SetId(id string) *Form {
	f.id = id
	return f
}

func (f *Form) SetParam(key, value string) *Form {
	f.params[key] = value
	return f
}

func (f *Form) DeleteParam(key string) *Form {
	delete(f.params, key)
	return f
}

func (f *Form) AddCss(key, value string) *Form {
	f.css[key] = value
	return f
}

func (f *Form) RemoveCss(key string) *Form {
	delete(f.css, key)
	return f
}
