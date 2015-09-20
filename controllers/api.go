package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"../models"
	"../utils"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type apiMarkdownMember struct {
	Markdown string
}

func apiMarkdownHandler(document http.ResponseWriter, request *http.Request) {

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

	document.Header().Set("Content-Type", "application/json")
	document.Write(jso)
}

func apiFileUploadHandler(document http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(32 << 20)

	pageName, err := url.QueryUnescape(request.FormValue("page"))
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		return
	}

	var page = new(models.Page)

	err = page.LoadFromTitle(pageName)
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		return
	}

	var attachment models.File
	var buffer = new(bytes.Buffer)

	file, handler, err := request.FormFile("file")
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		return
	}

	defer file.Close()

	io.Copy(buffer, file)

	name, _ := url.QueryUnescape(handler.Filename)
	attachment.Name = name
	attachment.Type = handler.Header.Get("Content-Type")
	attachment.Data = buffer.Bytes()

	page.Attachments.Files = append(page.Attachments.Files, attachment)

	err = page.Update(page.Id)
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		return
	}

	document.WriteHeader(200)
}

func apiFileViewHandler(document http.ResponseWriter, request *http.Request) {

	var pageName = strings.Split(request.URL.Path, "/")[3]
	var fileName = strings.Split(request.URL.Path, "/")[4]

	var page = new(models.Page)

	err := page.LoadFromTitle(pageName)
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		document.WriteHeader(404)
		return
	}

	for _, file := range page.Attachments.Files {

		if file.Name == fileName {

			document.Header().Set("Content-Type", file.Type)

			document.Write(file.Data)

			return
		}
	}

	utils.PromulgateDebugStr(os.Stdout, "ImageFile "+fileName+" of "+pageName+" NotFound")

	document.WriteHeader(404)
	return
}
