package controllers

import (
	"html"
	"net/http"
	"os"

	"go-wiki/models"
	"go-wiki/templates"
	"go-wiki/utils"
)

type userMember struct {
	*templates.DefaultMember
	Information models.User
}

func userViewHandler(document http.ResponseWriter, request *http.Request) {

	var requestedUser = html.EscapeString(request.URL.Path[11:])
	var user models.User
	var tmpl templates.Template

	err := user.Load(requestedUser)
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		panic(err.Error())
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
		panic(err.Error())
	}

}

func userEditHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "userEdit.tmpl"

}

func userRegisterHandler(document http.ResponseWriter, request *http.Request) {

	var tmpl templates.Template

	tmpl.Layout = "default.tmpl"
	tmpl.Template = "userRegister.tmpl"

	if request.Method == "POST" {

		var user models.User
		user.Name = request.FormValue("Name")
		user.Address = request.FormValue("Address")
		user.Password = models.GenerateHash(request.FormValue("Password"))

		err := user.Save(user.Name)
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			panic(err.Error())
		}

		tmpl.Template = "userRegisterSuccess.tmpl"
		tmpl.Render(document, &templates.DefaultMember{
			Title: "登録完了",
			User:  getSessionUser(request),
		})

	} else {

		tmpl.Render(document, &templates.DefaultMember{
			Title: "新規登録",
			User:  getSessionUser(request),
		})

	}
}

func userLoginHandler(document http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {

		var user models.User

		err := user.Load(request.FormValue("Name"))
		if err != nil {
			utils.PromulgateDebug(os.Stdout, err)
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
			http.Redirect(document, request, "/user/login#error", http.StatusFound)
			return
		}

	} else {
		var tmpl templates.Template

		tmpl.Layout = "default.tmpl"
		tmpl.Template = "userLogin.tmpl"

		err := tmpl.Render(document, &templates.DefaultMember{
			Title: "ログイン",
			User:  getSessionUser(request),
		})
		if err != nil {
			utils.PromulgateFatal(os.Stdout, err)
			panic(err.Error())
		}
	}
}

func userLogoutHandler(document http.ResponseWriter, request *http.Request) {

	session, _ := sessionStore.Get(request, "go-wiki")

	session.Values["User"] = "anonymous"

	session.Save(request, document)

	http.Redirect(document, request, request.Referer(), http.StatusFound)
}
