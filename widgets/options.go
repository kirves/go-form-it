package widgets

import (
	"fmt"
	"html/template"
)

func RadioButton(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/options/radiobutton.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func SelectMenu(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/options/select.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func Checkbox(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/options/checkbox.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
