package forms

import (
	"github.com/kirves/revel-forms/common"
	"github.com/kirves/revel-forms/fields"
	"testing"
)

var (
	txt, psw, btn fields.FieldInterface
)

func TestFieldRender(t *testing.T) {
	field := fields.TextField("test")
	field.AddClass("test")
	field.AddClass("class")
	field.SetId("testId")
	field.SetParam("param1", "val1")
	field.AddCss("css1", "val1")
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered field:", field.Render())
	txt = field
}

func TestPasswordRender(t *testing.T) {
	field := fields.PasswordField("test")
	field.AddClass("test")
	field.AddClass("class")
	field.SetId("testId")
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered field:", field.Render())
	psw = field
}

func TestButtonRender(t *testing.T) {
	field := fields.SubmitButton("btn", "Click me!")
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered button:", field.Render())
	btn = field
}

func TestRadioButtonRender(t *testing.T) {
	field := fields.RadioField("radio", map[string]string{
		"choice1": "value1",
		"choice2": "value2",
	})
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered radio:", field.Render())
}

func TestSelectRender(t *testing.T) {
	field := fields.SelectField("select", map[string]string{
		"choice1": "value1",
		"choice2": "value2",
	})
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered select:", field.Render())
}

func TestHiddenRender(t *testing.T) {
	field := fields.HiddenField("hidden")
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered hidden:", field.Render())
}

func TestNumberRender(t *testing.T) {
	field := fields.NumberField("number")
	field.SetStyle(formcommon.BASE)
	t.Log("Rendered number:", field.Render())
}

func TestFormRender(t *testing.T) {
	form := BaseForm(POST, "")
	form.AddField(txt)
	form.AddField(psw)
	form.AddField(btn)
	t.Log("Rendered form:", form.Render())
}

func TestFormFromModel(t *testing.T) {
	type Model struct {
		User     string
		Password string `form_widget:"password"`
	}

	form := BaseFormFromModel(Model{}, POST, "")
	t.Log("Rendered form:", form.Render())
}
