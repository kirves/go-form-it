package fields

import (
	"fmt"
	"github.com/kirves/go-form-it/common"
	"reflect"
)

// // Number field type.
// type NumberType struct {
// 	Field
// }

// // Range field type.
// type RangeType struct {
// 	Field
// }

// RangeField creates a default range field with the provided name. Min, max and step parameters define the expected behavior
// of the HTML field.
func RangeField(name string, min, max, step int) *Field {
	ret := FieldWithType(name, formcommon.RANGE)
	ret.SetParam("min", string(min))
	ret.SetParam("max", string(max))
	ret.SetParam("step", string(step))
	return ret
}

// NumberField craetes a default number field with the provided name.
func NumberField(name string) *Field {
	ret := FieldWithType(name, formcommon.NUMBER)
	return ret
}

// NumberFieldFromInstance creates and initializes a number field based on its name, the reference object instance and field number.
// This method looks for "form_min", "form_max" and "form_value" tags to add additional parameters to the field.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func NumberFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := NumberField(name)
	// check tags
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	if v := t.Get("form_min"); v != "" {
		ret.SetParam("min", v)
	}
	if v := t.Get("form_max"); v != "" {
		ret.SetParam("max", v)
	}
	if v := t.Get("form_value"); v != "" {
		ret.SetValue(v)
	} else {
		ret.SetValue(fmt.Sprintf("%v", reflect.ValueOf(i).Field(fieldNo).Interface()))
	}
	return ret
}

// RangeFieldFromInstance creates and initializes a range field based on its name, the reference object instance and field number.
// This method looks for "form_min", "form_max", "form_step" and "form_value" tags to add additional parameters to the field.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func RangeFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := RangeField(name, 0, 10, 1)
	// check tags
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	if v := t.Get("form_min"); v != "" {
		ret.SetParam("min", v)
	}
	if v := t.Get("form_max"); v != "" {
		ret.SetParam("max", v)
	}
	if v := t.Get("form_step"); v != "" {
		ret.SetParam("step", v)
	}
	if v := t.Get("form_value"); v != "" {
		ret.SetValue(v)
	} else {
		ret.SetValue(fmt.Sprintf("%v", reflect.ValueOf(i).Field(fieldNo).Interface()))
	}
	return ret
}
