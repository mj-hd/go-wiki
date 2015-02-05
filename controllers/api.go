package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"go-wiki/models"
	"go-wiki/utils"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
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

func apiFileUploadHandler(document http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(32 << 20)

	var pageName = request.FormValue("page")

	var page = new(models.Page)
	err := page.Load(pageName)

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

	err = page.Save(pageName)
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
	err := page.Load(pageName)

	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		// エラー画像
		return
	}

	for _, file := range page.Attachments.Files {

		if file.Name == fileName {

			document.Header().Set("Content-Type", file.Type)

			document.Write(file.Data)

		}
	}

	utils.PromulgateDebugStr(os.Stdout, "ImageFile "+fileName+" of "+pageName+" NotFound")

	// エラー画像
	return
}
