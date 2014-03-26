package widgets

import (
	"fmt"
	"html/template"
)

func Number(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/number/number.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func Range(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/number/range.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
