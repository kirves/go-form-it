package fields

import (
	"fmt"
	"github.com/kirves/revel-forms/common"
	"reflect"
)

type NumberType struct {
	Field
}

type RangeType struct {
	Field
}

func RangeField(name string, min, max, step int) *RangeType {
	ret := &RangeType{
		FieldWithType(name, formcommon.RANGE),
	}
	ret.SetParam("min", string(min))
	ret.SetParam("max", string(max))
	ret.SetParam("step", string(step))
	return ret
}

func NumberField(name string) *NumberType {
	ret := &NumberType{
		FieldWithType(name, formcommon.NUMBER),
	}
	return ret
}

func NumberFieldFromInstance(i interface{}, fieldNo int, name string) *NumberType {
	ret := &NumberType{
		FieldWithType(name, formcommon.NUMBER),
	}
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
		ret.SetValue(fmt.Sprintf("%d", reflect.ValueOf(i).Field(fieldNo).Interface()))
	}
	return ret
}

func RangeFieldFromInstance(i interface{}, fieldNo int, name string) *RangeType {
	ret := &RangeType{
		FieldWithType(name, formcommon.NUMBER),
	}
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
		ret.SetValue(fmt.Sprintf("%d", reflect.ValueOf(i).Field(fieldNo).Interface()))
	}
	return ret
}
