package plugins

import (
	"html/template"
	"os"

	"go-wiki/models"
	"go-wiki/utils"
)

func Sidebar() template.HTML {

	var page models.Page

	err := page.Load("_sidebar")
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		panic(err.Error())
	}

	return template.HTML(string(page.Markdown()))
}
