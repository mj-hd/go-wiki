package models

import (
	"os"
	"time"

	"go-wiki/utils"

	"code.google.com/p/go.crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int
	Name       string
	Caption    string
	Level      int
	Registered time.Time
	Address    string
	Password   string
}

func (this *User) Load(name string) error {

	row := DB.QueryRow("SELECT id, name, caption, level, registered, address, password FROM users WHERE name = ?", name)
	err := row.Scan(&this.Id, &this.Name, &this.Caption, &this.Level, &this.Registered, &this.Address, &this.Password)

	return err
}

func (this *User) Save(name string) error {

	if name == "" || !UserExists(name) {
		_, err := DB.Exec("INSERT INTO users ( name, caption, level, registered, address, password ) VALUES ( ?, ?, ?, ?, ?, ?)", this.Name, this.Caption, this.Level, this.Registered, this.Address, this.Password)
		if err != nil {
			return err
		}
	} else {
		_, err := DB.Exec("UPDATE users SET name = ?, caption = ?, level = ?, registered = ?, address = ?, password = ? WHERE name = ?", this.Name, this.Caption, this.Level, this.Registered, this.Address, this.Password, name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *User) Login(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(pass)) == nil
}

func GenerateHash(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		panic(err.Error())
	}
	return string(hash)
}

func UserExists(name string) bool {

	var rowCount int

	err := DB.QueryRow("SELECT count(id) FROM users WHERE title = ?", name).Scan(&rowCount)
	if err != nil {
		return false
	}

	return rowCount > 0
}
