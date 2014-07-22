package controllers

import (
	"html"
	"html/template"
	"net/http"

	"go-wiki/models"
	"go-wiki/templates"
)

type pageMember struct {
	Title    string
	Markdown template.HTML
}
type pageListMember struct {
	Title string
	Pages []models.Page
}

// Markdownを表示する
func pageHandler(document http.ResponseWriter, request *http.Request) {

	requestedPage := html.EscapeString(request.URL.Path[6:])

	if requestedPage == "" {

		var pages []models.Page

		err := models.GetPageList(&pages, 10)
		if err != nil {
			panic(err.Error())
		}

		var tmpl templates.Template
		tmpl.Layout = "default.tmpl"
		tmpl.Template = "pagelist.tmpl"
		tmpl.Render(document, pageListMember{Title: "ページ一覧", Pages: pages})
	} else {
		var page models.Page
		err := page.Load(requestedPage)
		if err != nil {
			panic(err.Error())
		}

		var tmpl templates.Template

		tmpl.Layout = "default.tmpl"
		tmpl.Template = "page.tmpl"

		tmpl.Render(document, pageMember{Title: page.Title, Markdown: template.HTML(page.Markdown())})
	}

}

// Markdownを編集する
func pageEditHandler(document http.ResponseWriter, request *http.Request) {

}

// Markdownを作成する
func pageCreateHandler(document http.ResponseWriter, request *http.Request) {

}
