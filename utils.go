package forms

import (
	"bytes"
	"github.com/kirves/go-form-it/fields"
	"html/template"
	"reflect"
)

type FormElement interface {
	Render() template.HTML
}

func (f *Form) Render() template.HTML {
	var s string
	buf := bytes.NewBufferString(s)
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

func (f *Form) Elements(elems ...FormElement) *Form {
	for _, e := range elems {
		t := reflect.TypeOf(e)
		switch {
		case t.Implements(reflect.TypeOf((*fields.FieldInterface)(nil)).Elem()):
			f.addField(e.(fields.FieldInterface))
		case reflect.ValueOf(e).Type().String() == "*forms.FieldSetType":
			f.addFieldSet(e.(*FieldSetType))
		}
	}
	return f
}

func (f *Form) addField(field fields.FieldInterface) *Form {
	field.SetStyle(f.style)
	f.fields = append(f.fields, field)
	f.fieldMap[field.Name()] = len(f.fields) - 1
	return f
}

func (f *Form) addFieldSet(fs *FieldSetType) *Form {
	for _, v := range fs.fields {
		v.SetStyle(f.style)
	}
	f.fields = append(f.fields, fs)
	f.fieldMap[fs.Name()] = len(f.fields) - 1
	return f
}

func (f *Form) RemoveField(name string) *Form {
	ind, ok := f.fieldMap[name]
	if !ok {
		return f
	}
	delete(f.fieldMap, name)
	f.fields = append(f.fields[:ind], f.fields[ind+1:]...)
	return f
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
