package main

import (
	"net/http"

	"go-wiki/config"
	"go-wiki/controllers"
	"go-wiki/models"
)

func main() {

	defer models.Del()
	defer controllers.Del()

	for path, function := range controllers.Routes {
		http.HandleFunc(path, function)
	}

	http.Handle("/"+config.StaticPath, http.StripPrefix("/static/", http.FileServer(http.Dir(config.StaticPath))))

	http.ListenAndServe(":"+config.ServerPort, nil)
}
