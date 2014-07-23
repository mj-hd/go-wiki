package controllers

import (
	"html/template"
	"net/http"
	"os"

	"go-wiki/models"
	"go-wiki/templates"
	"go-wiki/utils"
)

func indexHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template
	var page models.Page

	err := page.Load("index")
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		panic(err.Error())
	}

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "pageView.tmpl"

	tmpl.Render(document, pageMember{
		DefaultMember: &templates.DefaultMember{
			Title: "index page",
		},
		Markdown: template.HTML(page.Markdown()),
	})
}
