package main

import (
	"net/http"
	"os"

	"go-wiki/config"
	"go-wiki/controllers"
	"go-wiki/models"
	"go-wiki/plugins"
	"go-wiki/templates"
	"go-wiki/utils"

	"github.com/gorilla/context"
)

func main() {

	utils.LogFile = config.LogFile
	utils.DisplayLog = config.DisplayLog
	utils.LogLevel = config.LogLevel

	utils.PromulgateDebugStr(os.Stdout, "初期化を開始...")

	defer models.Del()
	defer controllers.Del()
	defer templates.Del()
	defer plugins.Del()

	for route := range controllers.Router.Iterator() {
		http.HandleFunc(route.Path, route.Function)
		utils.PromulgateDebugStr(os.Stdout, route.Path+"に関数を割当")
	}

	http.Handle("/"+config.StaticPath, http.StripPrefix("/"+config.StaticPath, http.FileServer(http.Dir(config.StaticPath))))
	utils.PromulgateDebugStr(os.Stdout, "/"+config.StaticPath+"に静的コンテンツを割当")

	utils.PromulgateInfoStr(os.Stdout, "ポート"+config.ServerPort+"でサーバを開始...")
	http.ListenAndServe(":"+config.ServerPort, context.ClearHandler(http.DefaultServeMux))
}
