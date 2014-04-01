package widgets

import (
	"bytes"
	"fmt"
	"github.com/kirves/go-form-it/common"
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
	w.Template.ExecuteTemplate(buf, "main", data)
	return buf.String()
}

func BaseWidget(style, inputType string) *Widget {
	var urls []string = []string{"templates/%s/generic.tmpl"}
	switch inputType {
	case formcommon.BUTTON:
		urls = append(urls, "templates/%s/button.html")
	case formcommon.CHECKBOX:
		urls = append(urls, "templates/%s/options/checkbox.html")
	case formcommon.TEXTAREA:
		urls = append(urls, "templates/%s/text/textareainput.html")
	case formcommon.SELECT:
		urls = append(urls, "templates/%s/options/select.html")
	case formcommon.PASSWORD:
		urls = append(urls, "templates/%s/text/passwordinput.html")
	case formcommon.RADIO:
		urls = append(urls, "templates/%s/options/radiobutton.html")
	case formcommon.TEXT:
		urls = append(urls, "templates/%s/text/textinput.html")
	case formcommon.RANGE:
		urls = append(urls, "templates/%s/number/range.html")
	case formcommon.NUMBER:
		urls = append(urls, "templates/%s/number/number.html")
	case formcommon.RESET:
		urls = append(urls, "templates/%s/button.html")
	case formcommon.SUBMIT:
		urls = append(urls, "templates/%s/button.html")
	case formcommon.DATE:
		urls = append(urls, "templates/%s/datetime/date.html")
	case formcommon.DATETIME:
		urls = append(urls, "templates/%s/datetime/datetime.html")
	case formcommon.TIME:
		urls = append(urls, "templates/%s/datetime/time.html")
	case formcommon.DATETIME_LOCAL:
		urls = append(urls, "templates/%s/datetime/datetime.html")
	case formcommon.STATIC:
		urls = append(urls, "templates/%s/static.html")
	case formcommon.SEARCH,
		formcommon.TEL,
		formcommon.URL,
		formcommon.WEEK,
		formcommon.COLOR,
		formcommon.EMAIL,
		formcommon.FILE,
		formcommon.HIDDEN,
		formcommon.IMAGE,
		formcommon.MONTH:
		urls = append(urls, "templates/%s/input.html")
	default:
		urls = append(urls, "templates/%s/input.html")
	}
	styledUrls := make([]string, len(urls))
	for i := range urls {
		styledUrls[i] = fmt.Sprintf(urls[i], style)
	}
	templ, err := template.ParseFiles(styledUrls...)
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}

func GenericWidget(style string) *Widget {
	templ, err := template.ParseFiles(fmt.Sprintf("templates/%s/input.html", style))
	if err != nil {
		panic(err)
	}
	return &Widget{templ}
}
