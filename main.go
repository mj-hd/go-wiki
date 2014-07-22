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
	http.ListenAndServe(":"+config.ServerPort, nil)
}
