package widgets

import (
	"fmt"
	"html/template"
)

func TextInput(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/text/textinput.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func PasswordInput(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/text/passwordinput.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func TextAreaInput(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/text/textareainput.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
