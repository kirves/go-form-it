package fields

import (
	"github.com/kirves/revel-forms/common"
)

type DatetimeType struct {
	Field
}

type DateType struct {
	Field
}

type TimeType struct {
	Field
}

func DatetimeField(name string) *DatetimeType {
	ret := &DatetimeType{
		FieldWithType(name, formcommon.DATETIME),
	}
	return ret
}

func DateField(name string) *DateType {
	ret := &DateType{
		FieldWithType(name, formcommon.DATE),
	}
	return ret
}

func TimeField(name string) *TimeType {
	ret := &TimeType{
		FieldWithType(name, formcommon.TIME),
	}
	return ret
}
