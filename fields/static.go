package fields

import (
	"fmt"
	"github.com/kirves/go-form-it/common"
	"reflect"
)

type StaticType struct {
	Field
}

func StaticField(name, content string) *StaticType {
	ret := &StaticType{
		FieldWithType(name, formcommon.STATIC),
	}
	ret.SetText(content)
	return ret
}

func StaticFieldFromInstance(i interface{}, fieldNo int, name string) *StaticType {
	ret := &StaticType{
		FieldWithType(name, formcommon.STATIC),
	}
	ret.SetText(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).Interface()))
	return ret
}
