package controllers

import (
	"html/template"
	"net/http"
	"os"

	"go-wiki/config"
	"go-wiki/models"
	"go-wiki/templates"
	"go-wiki/utils"
)

func indexHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template
	var page models.Page

	err := page.LoadFromTitle("index")
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		showError(document, request, "インデックスページが存在しません。管理人に問い合わせてください。")
		return
	}

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "pageView.tmpl"

	err = tmpl.Render(document, pageMember{
		DefaultMember: &templates.DefaultMember{
			Title: config.SiteTitle,
			User:  getSessionUser(request),
		},
		Markdown: template.HTML(page.Markdown()),
	})
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		showError(document, request, "ページの表示に失敗しました。")
		return
	}
}
