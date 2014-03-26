package fields

import (
	"github.com/kirves/revel-forms/common"
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
