package templates

import (
	"fmt"
	"html/template"
	"io"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w io.Writer, data interface{}) error {
	return t.htmlTpl.Execute(w, data)
}

func ParseFS(pattern ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(FS, pattern...)
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
