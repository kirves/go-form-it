package forms

import (
	"github.com/kirves/revel-forms/common"
	"github.com/kirves/revel-forms/fields"
	"testing"
	"time"
)

const (
	style = formcommon.BOOTSTRAP
)

var (
	txt, psw, btn fields.FieldInterface
)

func TestFieldRender(t *testing.T) {
	field := fields.TextField("test")
	field.AddClass("test").AddClass("class").SetId("testId").SetParam("param1", "val1").AddCss("css1", "val1").SetStyle(formcommon.BASE).Disabled()
	field.AddLabelClass("LABEL")
	field.SetLabel("This is a label")
	t.Log("Rendered field:", field.Render())
	txt = field
}

func TestPasswordRender(t *testing.T) {
	field := fields.PasswordField("test")
	field.AddClass("test")
	field.AddClass("class")
	field.SetId("testId")
	field.SetStyle(style)
	t.Log("Rendered field:", field.Render())
	psw = field
}

func TestButtonRender(t *testing.T) {
	field := fields.SubmitButton("btn", "Click me!")
	field.SetStyle(style)
	t.Log("Rendered button:", field.Render())
	btn = field
}

func TestRadioButtonRender(t *testing.T) {
	field := fields.RadioField("radio", map[string]string{
		"choice1": "value1",
		"choice2": "value2",
	})
	field.SetStyle(style)
	t.Log("Rendered radio:", field.Render())
}

func TestSelectRender(t *testing.T) {
	field := fields.SelectField("select", map[string]string{
		"choice1": "value1",
		"choice2": "value2",
	})
	field.SetStyle(style)
	t.Log("Rendered select:", field.Render())
}

func TestHiddenRender(t *testing.T) {
	field := fields.HiddenField("hidden")
	field.SetStyle(style)
	t.Log("Rendered hidden:", field.Render())
}

func TestNumberRender(t *testing.T) {
	field := fields.NumberField("number")
	field.SetStyle(style)
	t.Log("Rendered number:", field.Render())
}

func TestFormRender(t *testing.T) {
	form := BootstrapForm(POST, "")
	form.AddField(txt)
	form.AddField(psw)
	form.AddField(btn)
	t.Log("Rendered form:", form.Render())
}

func TestFormFromModel(t *testing.T) {
	type Model struct {
		User      string    `form_label:"User label test"`
		password  string    `form_widget:"password"`
		Id        int       `form_min:"0" form_max:"5"`
		Ts        time.Time `form_min:"2013-04-22T15:00"`
		RadioTest string    `form_widget:"select" form_choices:"A|Option A|B|Option B" form_value:"A"`
		BoolTest  bool      //`form_options:"checked"`
	}

	form := BaseFormFromModel(Model{"asd", "lol", 20, time.Now(), "B", false}, POST, "")
	t.Log("Rendered form:", form.Render())
}
