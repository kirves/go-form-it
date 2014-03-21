package widgets

import (
	"bytes"
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
