package controllers

import (
	"html"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"

	"go-wiki/models"
	"go-wiki/templates"
	"go-wiki/utils"
	"go-wiki/config"
)

type userMember struct {
	*templates.DefaultMember
	Information models.User
}
type userLoginMember struct {
	*templates.DefaultMember
	Message string
	EnableRegister bool
}

func userViewHandler(document http.ResponseWriter, request *http.Request) {

	var requestedUser = html.EscapeString(request.URL.Path[11:])
	var user models.User
	var tmpl templates.Template

	err := user.Load(requestedUser)
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		showError(document, request, "ユーザが存在しません。")
		return
	}

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "userView.tmpl"

	err = tmpl.Render(document, userMember{
		DefaultMember: &templates.DefaultMember{
			Title: "ユーザ" + user.Name,
			User:  getSessionUser(request),
		},
		Information: user,
	})
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		showError(document, request, "ページの表示に失敗しました。")
		return
	}

}

func userEditHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "userEdit.tmpl"

	showError(document, request, "未作成のページです。")
}

func userRegisterHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template

	if !config.EnableRegister {

		showError(document, request, "ユーザの新規登録は禁止されています。")

		return
	}

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "userRegister.tmpl"

	if request.Method == "POST" {

		var user models.User
		user.Name = bluemonday.UGCPolicy().Sanitize(request.FormValue("Name"))
		user.Address = bluemonday.UGCPolicy().Sanitize(request.FormValue("Address"))
		user.Password = models.GenerateHash(request.FormValue("Password"))

		if user.Name == "" || user.Address == "" || user.Password == "" {
			utils.PromulgateDebugStr(os.Stdout, "登録に必要な情報が足りません。")
			showError(document, request, "登録に必要な情報が足りません。")
			return
		}

		err := user.Save(user.Name)
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "ユーザの保存に失敗しました。")
			return
		}

		session, _ := sessionStore.Get(request, "go-wiki")
		session.AddFlash("ユーザ登録に成功しました。")
		session.Save(request, document)
		http.Redirect(document, request, "/success/", http.StatusFound)
		return

	} else {

		err := tmpl.Render(document, &templates.DefaultMember{
			Title: "新規登録",
			User:  getSessionUser(request),
		})
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "ページの表示に失敗しました。")
			return
		}

	}
}

func userLoginHandler(document http.ResponseWriter, request *http.Request) {

	session, _ := sessionStore.Get(request, "go-wiki")

	if request.Method == "POST" {

		var user models.User

		err := user.Load(bluemonday.UGCPolicy().Sanitize(request.FormValue("Name")))
		if err != nil {
			utils.PromulgateDebug(os.Stdout, err)
			session.AddFlash("ユーザが存在しません。")
			session.Save(request, document)
			http.Redirect(document, request, "/user/login#error", http.StatusFound)
			return
		}

		if user.Login(request.FormValue("Password")) {
			session, err := sessionStore.Get(request, "go-wiki")
			session.Values["User"] = user.Name
			session.Save(request, document)
			if err != nil {
				utils.PromulgateDebug(os.Stdout, err)
			}
			http.Redirect(document, request, "/", http.StatusFound)
			return
		} else {
			utils.PromulgateDebugStr(os.Stdout, "パスワード不一致")
			session.AddFlash("パスワードが違います。")
			session.Save(request, document)
			http.Redirect(document, request, "/user/login#error", http.StatusFound)
			return
		}

	} else {
		var tmpl templates.Template

		tmpl.Layout = "default.tmpl"
		tmpl.Template = "userLogin.tmpl"

		var message string
		flashes := session.Flashes()
		session.Save(request, document)

		if len(flashes) >= 1 {
			message = flashes[0].(string)
		} else {
			message = ""
		}

		err := tmpl.Render(document, userLoginMember{
			DefaultMember: &templates.DefaultMember{
				Title: "ログイン",
				User:  getSessionUser(request),
			},
			Message: message,
			EnableRegister: config.EnableRegister,
		})
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			showError(document, request, "ページの表示に失敗しました。")
			return
		}
	}
}

func userLogoutHandler(document http.ResponseWriter, request *http.Request) {

	session, _ := sessionStore.Get(request, "go-wiki")

	session.Values["User"] = "anonymous"

	session.Save(request, document)

	http.Redirect(document, request, request.Referer(), http.StatusFound)
}
