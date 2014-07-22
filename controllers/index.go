package controllers

import (
	"html/template"
	"net/http"

	"go-wiki/models"
	"go-wiki/templates"
)

func indexHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template
	var page models.Page

	err := page.Load("index")
	if err != nil {
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
