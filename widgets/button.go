package widgets

import (
	"fmt"
	"html/template"
)

func Button(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/button.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
