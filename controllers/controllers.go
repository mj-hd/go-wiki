package controllers

import "net/http"

var Routes = map[string]func(http.ResponseWriter, *http.Request){
	"/":            indexHandler,
	"/page/":       pageHandler,
	"/page/edit":   pageEditHandler,
	"/page/create": pageCreateHandler,
}

func init() {

}
func Del() {

}
