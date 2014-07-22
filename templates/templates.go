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

type DefaultMember struct {
	Title string
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

func (this *DefaultMember) LinkCSS(cssFile string) template.HTML {
	return template.HTML("<link rel='stylesheet' href='/" + config.CssPath + cssFile + "' type='text/css' />")
}
func (this *DefaultMember) EmbedImage(imgFile string, alt string) template.HTML {
	return template.HTML("<img alt='" + alt + "' src='/" + config.ImgPath + imgFile + "' />")
}
func (this *DefaultMember) LinkJS(jsFile string) template.HTML {
	return template.HTML("<script type='text/javascript' src='/" + config.JsPath + jsFile + "' ></script>")
}
