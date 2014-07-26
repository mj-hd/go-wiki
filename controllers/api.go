package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/russross/blackfriday"
)

type apiMarkdownMember struct {
	Markdown string
}

func apiMarkdownHandler(document http.ResponseWriter, request *http.Request) {
	document.Header().Set("Content-Type", "application/json")
	var jso []byte

	jso, _ = json.Marshal(apiMarkdownMember{Markdown: string(blackfriday.MarkdownCommon([]byte(request.URL.Query()["text"][0])))})
	document.Write(jso)
}
