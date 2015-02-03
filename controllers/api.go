package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

type apiMarkdownMember struct {
	Markdown string
}

func apiMarkdownHandler(document http.ResponseWriter, request *http.Request) {
	document.Header().Set("Content-Type", "application/json")
	var jso []byte
	var text string
	var texts = request.URL.Query()["text"]

	if len(texts) == 0 {
		text = ""
	} else {
		text = texts[0]
	}

	jso, _ = json.Marshal(
		apiMarkdownMember{
			Markdown: string(
				bluemonday.UGCPolicy().SanitizeBytes(
					blackfriday.MarkdownCommon(
						[]byte(text))))})
	document.Write(jso)
}
