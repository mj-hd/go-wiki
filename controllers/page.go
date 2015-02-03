package controllers

import (
	"database/sql"
	"html"
	"html/template"
	"net/http"
	"os"
	"time"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"

	"go-wiki/models"
	"go-wiki/templates"
	"go-wiki/utils"
	"go-wiki/config"
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
			showError(document, request, "ページ一覧の取得に失敗しました。")
			return
		}

		tmpl.Template = "pageList.tmpl"
		err = tmpl.Render(document, pageListMember{
			DefaultMember: &templates.DefaultMember{
				Title: "ページ一覧",
				User:  getSessionUser(request),
			},
			Pages: pages,
		})
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "ページの表示に失敗しました。")
			return
		}

	} else {
		var page models.Page
		err := page.Load(requestedPage)
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "指定されたページは存在しません。")
			return
		}

		tmpl.Template = "pageView.tmpl"

		err = tmpl.Render(document, pageMember{
			DefaultMember: &templates.DefaultMember{
				Title: page.Title,
				User:  getSessionUser(request),
			},
			Markdown:    template.HTML(page.Markdown()),
			Information: page,
		})
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "ページの表示に失敗しました。")
			return
		}

	}

}

// Markdownを編集する
func pageEditHandler(document http.ResponseWriter, request *http.Request) {

	requestedPage := html.EscapeString(request.URL.Path[6+5:])

	if requestedPage == "" {
		requestedPage = "index"
	}

	var page models.Page
	if page.Load(requestedPage) == nil {

		var tmpl templates.Template
		user := getSessionUser(request)

		utils.PromulgateDebug(os.Stdout, errors.New("Trying to edit as " + user))
		if user == "anonymous" && !config.EnableAnonymousEdit {
			showError(document, request, "匿名の編集は禁止されています。")
			return
		}

		tmpl.Layout = "editor.tmpl"
		tmpl.Template = "pageEdit.tmpl"

		if page.Locked {
			if user != page.User {
				curUser := models.User{}
				pagUser := models.User{}

				curUser.Load(user)
				pagUser.Load(page.User)

				if curUser.Level >= pagUser.Level {
					showError(document, request, "ページはロックされています。")
					return
				}
			}
		}

		err := tmpl.Render(document, pageMember{
			DefaultMember: &templates.DefaultMember{
				Title: requestedPage + "の編集",
				User:  getSessionUser(request),
			},
			Markdown:    template.HTML(page.Markdown()),
			Information: page,
		})
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "ページの表示に失敗しました。")
			return
		}

	} else {
		showError(document, request, "存在しないページです。")
		return
	}
}

// Markdownを作成する
func pageCreateHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template
	var page models.Page

	if getSessionUser(request) == "anonymous" && !config.EnableAnonymousEdit {

		showError(document, request, "匿名の編集は禁止されています。ログインをしてください。")

		return
	}

	tmpl.Layout = "editor.tmpl"
	tmpl.Template = "pageEdit.tmpl"

	page.Title = "タイトル"
	page.User = getSessionUser(request)
	page.Locked = false
	page.Content = sql.NullString{String: "内容", Valid: true}

	err := tmpl.Render(document, pageMember{
		DefaultMember: &templates.DefaultMember{
			Title: "新規ページの作成",
			User:  page.User,
		},
		Markdown:    template.HTML(page.Markdown()),
		Information: page,
	})
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		showError(document, request, "ページの表示に失敗しました。")
		return
	}

}

func pageSaveHandler(document http.ResponseWriter, request *http.Request) {

	oldTitle := bluemonday.UGCPolicy().Sanitize(request.FormValue("OldTitle"))

	var page models.Page

	page.Title = bluemonday.UGCPolicy().Sanitize(request.FormValue("Title"))
	if page.Title == "" {
		showError(document, request, "タイトルが空です。")
		return
	}

	page.Content = sql.NullString{String: request.FormValue("Content"), Valid: true}
	page.User = getSessionUser(request)
	page.Locked = (request.FormValue("Locked") == "1")
	if !models.PageExists(page.Title) {
		page.Modified = mysql.NullTime{Valid: false}
		page.Created = time.Now()
	} else {
		if oldTitle != page.Title {
			showError(document, request, "同じタイトルのページが存在します。")
			return
		}

		page.Modified = mysql.NullTime{Time: time.Now(), Valid: true}
		var oldPage models.Page
		oldPage.Load(oldTitle)
		page.Created = oldPage.Created
	}

	if page.Save(oldTitle) == nil {

		var tmpl templates.Template
		tmpl.Layout = "default.tmpl"
		tmpl.Template = "pageSave.tmpl"

		http.Redirect(document, request, "/page/view/"+page.Title, http.StatusFound)

	} else {

		utils.PromulgateFatalStr(os.Stdout, "Page"+page.Title+"の保存に失敗")
		showError(document, request, "ページの保存に失敗しました。")
		return

	}

}
