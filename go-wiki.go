package main

import (
	"net/http"
	"os"
	"fmt"
	"syscall"

	"go-wiki/config"
	"go-wiki/controllers"
	"go-wiki/models"
	"go-wiki/plugins"
	"go-wiki/templates"
	"go-wiki/utils"

	"github.com/gorilla/context"
)

func main() {

	errcd := daemon(0, 0)

	if errcd != 0 {
		fmt.Println("デーモン化に失敗。")
		os.Exit(1)
	}

	os.Chdir(config.GowikiPath)

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

/*
daemon関数
*/
func daemon(nochdir, noclose int) int {
 var ret uintptr
 var err syscall.Errno
 
 // バックグラウンドプロセスにする
 // 子プロセスを生成し，親プロセスを終了する
 ret, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
 if err != 0 {
  return -1
 }
 switch ret {
  case 0:
   // 子プロセスが生成できたらそのまま処理を続ける
   break
  default:
   // 親プロセスだとここで終了する
   os.Exit(0)
 }
 
  
 // 新しいセッションを生成(子プロセスがセッションリーダになる)
 pid, _ := syscall.Setsid()
 if pid == -1 {
  return -1
 }
  
 if nochdir == 0 {
  // カレントディレクトリの再設定。ここではルートにしておく
  os.Chdir("/")
 }
  
 // ファイルのパーミッションを再設定(必須ではない。オプション)
 syscall.Umask(0)
 
 if noclose == 0 {
  // 標準入出力先を/dev/nullファイルに変更して、すべて破棄する
  f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
  if e == nil {
   fd := int(f.Fd())
   syscall.Dup2(fd, int(os.Stdin.Fd()))
   syscall.Dup2(fd, int(os.Stdout.Fd()))
   syscall.Dup2(fd, int(os.Stderr.Fd()))
  }
 }
 
 return 0
}
