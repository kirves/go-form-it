package forms

import (
	"github.com/kirves/go-form-it/common"
	"github.com/kirves/go-form-it/fields"
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
	field.AddClass("test").AddClass("class").SetId("testId").SetParam("param1", "val1").AddCss("css1", "val1").SetStyle(style).Disabled()
	field.AddLabelClass("LABEL")
	field.SetLabel("This is a label")
	field.AddError("ERROR")
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
	field := fields.RadioField("radio", []fields.InputChoice{
		fields.InputChoice{Id: "choice1", Val: "value1"},
		fields.InputChoice{Id: "choice2", Val: "value2"},
	})
	field.SetStyle(style)
	t.Log("Rendered radio:", field.Render())
}

func TestSelectRender(t *testing.T) {
	field := fields.SelectField("select", map[string][]fields.InputChoice{
		"": []fields.InputChoice{
			fields.InputChoice{"choice1", "value1"},
			fields.InputChoice{"choice2", "value2"},
		},
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
	form.Elements(&FieldSetType{}, txt, psw, btn)
	// form.AddField(psw)
	// form.AddField(btn)
	t.Log("Rendered form:", form.Render())
}

func TestFormFromSimpleModel(t *testing.T) {
	type User struct {
		Username  string
		Password1 string `form_widget:"password" form_label:"Password 1"`
		Password2 string `form_widget:"password" form_label:"Password 2"`
		SkipThis  int    `form_options:"skip"`
	}

	u := User{}

	form := BaseFormFromModel(u, POST, "/action.html")
	t.Log("Rendered form:", form.Render())
}

func TestFormFromModel(t *testing.T) {
	type Model struct {
		User      string    `form_label:"User label test"`
		password  string    `form_widget:"password"`
		Id        int       `form_min:"0" form_max:"5"`
		Ts        time.Time `form_min:"2013-04-22T15:00"`
		RadioTest string    `form_widget:"select" form_choices:"|A|Option A|G1|B|Option B" form_value:"A"`
		BoolTest  bool      //`form_options:"checked"`
	}

	form := BaseFormFromModel(Model{"asd", "lol", 20, time.Now(), "B", false}, POST, "")
	t.Log("Rendered form:", form.Render())
}

func TestBSFormFromModel(t *testing.T) {
	type Model struct {
		User      string    `form_label:"User label test"`
		password  string    `form_widget:"password"`
		Id        int       `form_min:"0" form_max:"5"`
		Ts        time.Time `form_min:"2013-04-22T15:00"`
		RadioTest string    `form_widget:"select" form_choices:"|A|Option A|G1|B|Option B" form_value:"A"`
		BoolTest  bool      //`form_options:"checked"`
	}

	form := BootstrapFormFromModel(Model{"asd", "lol", 20, time.Now(), "B", false}, POST, "")
	t.Log("Rendered form:", form.Render())
}

func TestInlineCreation(t *testing.T) {
	form := BaseForm(POST, "/action.html").Elements(
		fields.TextField("text_field").SetLabel("Username"),
		FieldSet("psw_fieldset",
			fields.PasswordField("psw1").AddClass("password_class").SetLabel("Password 1"),
			fields.PasswordField("psw2").AddClass("password_class").SetLabel("Password 2"),
		),
		fields.SubmitButton("btn1", "Submit"),
	)
	t.Log("Rendered form:", form.Render())
}
