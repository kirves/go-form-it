package widgets

import (
	"fmt"
	"html/template"
)

func RadioButton(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/radiobutton.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func SelectMenu(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/select.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
