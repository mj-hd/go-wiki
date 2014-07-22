package controllers

import "net/http"

var Routes = map[string]func(http.ResponseWriter, *http.Request){
	"/":             indexHandler,
	"/page/":        pageHandler,
	"/page/view/":   pageViewHandler,
	"/page/edit/":   pageEditHandler,
	"/page/create/": pageCreateHandler,
	"/page/save/":   pageSaveHandler,
}

func init() {

}
func Del() {

}
