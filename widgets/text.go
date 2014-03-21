package widgets

import (
	"fmt"
	"html/template"
)

func TextInput(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/textinput.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func PasswordInput(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/passwordinput.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func TextAreaInput(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/textareainput.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
