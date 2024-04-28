package templates

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w io.Writer, data interface{}) error {
	return t.htmlTpl.Execute(w, data)
}

func ParseFS(pattern ...string) (Template, error) {
	htmlTpl := template.New(pattern[0])
	htmlTpl = htmlTpl.Funcs(
		template.FuncMap{
			"jsonMarshal": func(v interface{}) template.JS {
				js, err := json.Marshal(v)
				if err != nil {
					log.Println("Error marshaling JSON:", err)
					return ""
				}
				return template.JS(js)
			},
		},
	)
	htmlTpl, err := htmlTpl.ParseFS(FS, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("error parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
