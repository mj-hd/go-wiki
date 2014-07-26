package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"

	"go-wiki/config"
)

var Router Routes
var sessionStore = sessions.NewCookieStore([]byte(config.SessionKey))

func init() {

	Router.Register("/", indexHandler)
	Router.Register("/page/", pageHandler)
	Router.Register("/page/view/", pageViewHandler)
	Router.Register("/page/edit/", pageEditHandler)
	Router.Register("/page/create/", pageCreateHandler)
	Router.Register("/page/save/", pageSaveHandler)
	Router.Register("/user/view/", userViewHandler)
	Router.Register("/user/edit/", userEditHandler)
	Router.Register("/user/register/", userRegisterHandler)
	Router.Register("/user/login/", userLoginHandler)
	Router.Register("/user/logout/", userLogoutHandler)
	Router.Register("/api/markdown/", apiMarkdownHandler)

}
func Del() {

}

func getSessionUser(request *http.Request) string {
	session, _ := sessionStore.Get(request, "go-wiki")
	if session.Values["User"] == nil {
		return "anonymous"
	}

	return session.Values["User"].(string)
}

type Routes struct {
	keys   []string
	values []func(http.ResponseWriter, *http.Request)
}
type Route struct {
	Path     string
	Function func(http.ResponseWriter, *http.Request)
}

func (this *Routes) Iterator() <-chan Route {
	ret := make(chan Route)

	go func() {
		for i, k := range this.keys {
			var route Route
			route.Path = k
			route.Function = this.values[i]

			ret <- route
		}
		close(ret)
	}()

	return ret
}

func (this *Routes) Register(path string, fn func(http.ResponseWriter, *http.Request)) {
	this.keys = append(this.keys, path)
	this.values = append(this.values, fn)
}

func (this *Routes) Value(path string) func(http.ResponseWriter, *http.Request) {
	for i, key := range this.keys {
		if path == key {
			return this.values[i]
		}
	}
	return nil
}

func (this *Routes) Key(fn *func(http.ResponseWriter, *http.Request)) string {
	for i, val := range this.values {
		if fn == &val {
			return this.keys[i]
		}
	}
	return ""
}
