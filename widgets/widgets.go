package widgets

import (
	"bytes"
	"fmt"
	"html/template"
)

type Widget struct {
	Template *template.Template
}

type WidgetInterface interface {
	Render(data interface{}) string
}

func (w *Widget) Render(data interface{}) string {
	var s string
	buf := bytes.NewBufferString(s)
	w.Template.Execute(buf, data)
	return buf.String()
}

func GenericWidget(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/input.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
