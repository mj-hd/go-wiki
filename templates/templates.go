package templates

import (
	"html/template"
	"io"

	"go-wiki/config"
)

type Template struct {
	Layout   string
	Template string
}

func init() {
}
func Del() {
}

func (this *Template) Render(w io.Writer, member interface{}) error {

	tmpl, err := template.ParseFiles(config.LayoutsPath+this.Layout, config.TemplatesPath+this.Template)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, member)
	if err != nil {
		return err
	}

	return nil
}
