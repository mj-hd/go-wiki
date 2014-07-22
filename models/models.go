package models

import (
	"database/sql"
	"fmt"

	"go-wiki/config"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err.Error())
	}
}

func Del() {
	DB.Close()
}
