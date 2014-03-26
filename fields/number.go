package fields

import (
	"github.com/kirves/revel-forms/common"
)

type NumberType struct {
	Field
}

func NumberField(name string) *NumberType {
	ret := &NumberType{
		FieldWithType(name, formcommon.NUMBER),
	}
	return ret
}
