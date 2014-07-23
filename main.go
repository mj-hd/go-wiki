package main

import (
	"net/http"
	"os"

	"go-wiki/config"
	"go-wiki/controllers"
	"go-wiki/models"
	"go-wiki/utils"
)

func main() {

	utils.LogFile = config.LogFile
	utils.DisplayLog = config.DisplayLog
	utils.LogLevel = config.LogLevel

	utils.PromulgateDebugStr(os.Stdout, "初期化を開始...")

	defer models.Del()
	defer controllers.Del()

	for path, function := range controllers.Routes {
		http.HandleFunc(path, function)
		utils.PromulgateDebugStr(os.Stdout, path+"に関数を割当")
	}

	http.Handle("/"+config.StaticPath, http.StripPrefix("/"+config.StaticPath, http.FileServer(http.Dir(config.StaticPath))))
	utils.PromulgateDebugStr(os.Stdout, "/"+config.StaticPath+"に静的コンテンツを割当")

	utils.PromulgateInfoStr(os.Stdout, "ポート"+config.ServerPort+"でサーバを開始...")
	http.ListenAndServe(":"+config.ServerPort, nil)
}
