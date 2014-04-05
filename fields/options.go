package fields

import (
	"fmt"
	"github.com/kirves/go-form-it/common"
	"reflect"
	"strings"
)

// Id - Value pair used to define an option for select and redio input fields.
type InputChoice struct {
	Id, Val string
}

// Radio button type.
type RadioType struct {
	Field
}

// Select field type.
type SelectType struct {
	Field
}

// Checkbox field type.
type CheckBoxType struct {
	Field
}

// =============== RADIO

// RadioField creates a default radio button input field with the provided name and list of choices.
func RadioField(name string, choices []InputChoice) *RadioType {
	ret := &RadioType{
		FieldWithType(name, formcommon.RADIO),
	}
	ret.additionalData["choices"] = []InputChoice{}
	ret.SetChoices(choices)
	return ret
}

// SetChoices takes as input a dictionary whose key-value entries are defined as follows: key is the group name (the empty string
// is the default group that is not explicitly rendered) and value is the list of choices belonging to that group.
// Grouping is only useful for Select fields, while groups are ignored in Radio fields.
func (f *RadioType) SetChoices(choices []InputChoice) *RadioType {
	f.additionalData["choices"] = choices
	return f
}

// RadioFieldFromInstance creates and initializes a radio field based on its name, the reference object instance and field number.
// This method looks for "form_choices" and "form_value" tags to add additional parameters to the field. "form_choices" tag is a list
// of <id>|<value> options, joined by "|" character; ex: "A|Option A|B|Option B" translates into 2 options: <A, Option A> and <B, Option B>.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func RadioFieldFromInstance(i interface{}, fieldNo int, name string) *RadioType {
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	choices := strings.Split(t.Get("form_choices"), "|")
	chArr := make([]InputChoice, 0)
	chMap := make(map[string]string)
	for i := 0; i < len(choices)-1; i += 2 {
		chArr = append(chArr, InputChoice{choices[i], choices[i+1]})
		chMap[choices[i]] = choices[i+1]
	}
	ret := RadioField(name, chArr)

	var v string = t.Get("form_value")
	if v == "" {
		v = fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String())
	}
	if _, ok := chMap[v]; ok {
		ret.SetValue(v)
	}
	return ret
}

// ================ SELECT

// SelectField creates a default select input field with the provided name and map of choices. Choices for SelectField are grouped
// by name (if <optgroup> is needed); "" group is the default one and does not trigger a <optgroup></optgroup> rendering.
func SelectField(name string, choices map[string][]InputChoice) *SelectType {
	ret := &SelectType{
		FieldWithType(name, formcommon.SELECT),
	}
	ret.additionalData["choices"] = map[string][]InputChoice{}
	ret.additionalData["multValues"] = map[string]struct{}{}
	ret.SetChoices(choices)
	return ret
}

// MultipleChoice configures the SelectField to accept and display multiple choices.
func (sf *SelectType) MultipleChoice() *SelectType {
	sf.AddTag("multiple")
	return sf
}

// SingleChoice configures the SelectField to accept and display only one choice.
func (sf *SelectType) SingleChoice() *SelectType {
	sf.RemoveTag("multiple")
	return sf
}

// If the SelectField is configured as "multiple", AddSelected adds a selected value to the field.
func (sf *SelectType) AddSelected(opt ...string) *SelectType {
	for _, v := range opt {
		sf.additionalData["multValues"].(map[string]struct{})[v] = struct{}{}
	}
	return sf
}

// If the SelectField is configured as "multiple", AddSelected removes the selected value from the field.
func (sf *SelectType) RemoveSelected(opt string) *SelectType {
	delete(sf.additionalData["multValues"].(map[string]struct{}), opt)
	return sf
}

// SetChoices takes as input a dictionary whose key-value entries are defined as follows: key is the group name (the empty string
// is the default group that is not explicitly rendered) and value is the list of choices belonging to that group.
// Grouping is only useful for Select fields, while groups are ignored in Radio fields.
func (sf *SelectType) SetChoices(choices map[string][]InputChoice) *SelectType {
	sf.additionalData["choices"] = choices
	return sf
}

// SelectFieldFromInstance creates and initializes a select field based on its name, the reference object instance and field number.
// This method looks for "form_choices" and "form_value" tags to add additional parameters to the field. "form_choices" tag is a list
// of <group<|<id>|<value> options, joined by "|" character; ex: "G1|A|Option A|G1|B|Option B" translates into 2 options in the same group G1:
// <A, Option A> and <B, Option B>. "" group is the default one.
// It also uses i object's [fieldNo]-th field content (if any) to override the "form_value" option and fill the HTML field.
func SelectFieldFromInstance(i interface{}, fieldNo int, name string, options map[string]struct{}) *SelectType {
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	choices := strings.Split(t.Get("form_choices"), "|")
	chArr := make(map[string][]InputChoice)
	chMap := make(map[string]string)
	for i := 0; i < len(choices)-2; i += 3 {
		if _, ok := chArr[choices[i]]; !ok {
			chArr[choices[i]] = make([]InputChoice, 0)
		}
		chArr[choices[i]] = append(chArr[choices[i]], InputChoice{choices[i+1], choices[i+2]})
		chMap[choices[i+1]] = choices[i+2]
	}
	ret := SelectField(name, chArr)

	if _, ok := options["multiple"]; ok {
		ret.MultipleChoice()
	}

	var v string = fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String())
	if v == "" {
		// TODO: multiple value parsing
		v = t.Get("form_value")
	}
	if _, ok := chMap[v]; ok {
		ret.SetValue(v)
	}
	return ret
}

// ================== CHECKBOX

// Checkbox creates a default checkbox field with the provided name. It also makes it checked by default based
// on the checked parameter.
func Checkbox(name string, checked bool) *CheckBoxType {
	ret := &CheckBoxType{
		FieldWithType(name, formcommon.CHECKBOX),
	}
	if checked {
		ret.AddTag("checked")
	}
	return ret
}

// CheckboxFromInstance creates and initializes a checkbox field based on its name, the reference object instance, field number and field options.
// It uses i object's [fieldNo]-th field content (if any) to override the "checked" option in the options map and check the field.
func CheckboxFromInstance(i interface{}, fieldNo int, name string, options map[string]struct{}) *CheckBoxType {
	ret := &CheckBoxType{
		FieldWithType(name, formcommon.CHECKBOX),
	}

	if _, ok := options["checked"]; ok {
		ret.AddTag("checked")
	} else {
		val := reflect.ValueOf(i).Field(fieldNo).Bool()
		if val {
			ret.AddTag("checked")
		}
	}
	return ret
}
