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
		utils.PromulgateFatalStr(os.Stdout, "サイドバープラグインが有効だが、_sidebarページが存在しない")
		panic(err.Error())
	}

	return template.HTML(string(page.Markdown()))
}
