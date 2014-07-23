package controllers

import (
	"database/sql"
	"html"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	"go-wiki/models"
	"go-wiki/templates"
	"go-wiki/utils"
)

type pageMember struct {
	*templates.DefaultMember
	Markdown    template.HTML
	Information models.Page
}
type pageListMember struct {
	*templates.DefaultMember
	Pages []models.Page
}

func pageHandler(document http.ResponseWriter, request *http.Request) {

	http.Redirect(document, request, "/page/view/", http.StatusFound)

}

// Markdownを表示する
func pageViewHandler(document http.ResponseWriter, request *http.Request) {

	requestedPage := html.EscapeString(request.URL.Path[6+5:])

	var tmpl templates.Template
	tmpl.Layout = "default.tmpl"

	if requestedPage == "" {

		var pages []models.Page

		err := models.GetPageList(&pages, 10)
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			panic(err.Error())
		}

		tmpl.Template = "pageList.tmpl"
		tmpl.Render(document, pageListMember{
			DefaultMember: &templates.DefaultMember{
				Title: "ページ一覧",
			},
			Pages: pages,
		})

	} else {
		var page models.Page
		err := page.Load(requestedPage)
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			panic(err.Error())
		}

		tmpl.Template = "pageView.tmpl"

		tmpl.Render(document, pageMember{
			DefaultMember: &templates.DefaultMember{
				Title: page.Title,
			},
			Markdown:    template.HTML(page.Markdown()),
			Information: page,
		})
	}

}

// Markdownを編集する
func pageEditHandler(document http.ResponseWriter, request *http.Request) {

	requestedPage := html.EscapeString(request.URL.Path[6+5:])

	if requestedPage == "" {
		http.Redirect(document, request, "/page/create", http.StatusFound)
		return
	} else {

		var page models.Page
		if page.Load(requestedPage) == nil {

			var tmpl templates.Template

			tmpl.Layout = "default.tmpl"
			tmpl.Template = "pageEdit.tmpl"

			tmpl.Render(document, pageMember{
				DefaultMember: &templates.DefaultMember{
					Title: requestedPage + "の編集",
				},
				Markdown:    template.HTML(page.Markdown()),
				Information: page,
			})

		} else {
			http.Redirect(document, request, "/page/create/"+requestedPage, http.StatusFound)
		}
	}
}

// Markdownを作成する
func pageCreateHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "pageCreate.tmpl"

	tmpl.Render(document, templates.DefaultMember{Title: "新規ページの作成"})
}

func pageSaveHandler(document http.ResponseWriter, request *http.Request) {

	oldTitle := request.FormValue("OldTitle")

	var page models.Page

	page.Title = request.FormValue("Title")
	page.Content = sql.NullString{String: request.FormValue("Content"), Valid: true}
	page.User = "admin"
	page.Locked = false
	page.Created = time.Now()
	if !models.PageExists(page.Title) {
		page.Modified = mysql.NullTime{Valid: false}
	} else {
		page.Modified = mysql.NullTime{Time: time.Now(), Valid: true}
	}

	if page.Save(oldTitle) == nil {

		var tmpl templates.Template
		tmpl.Layout = "default.tmpl"
		tmpl.Template = "pageSave.tmpl"

		tmpl.Render(document, nil)

	} else {

		utils.PromulgateFatalStr(os.Stdout, "Page"+page.Title+"の保存に失敗")
		panic("Save failed.")

	}

}
