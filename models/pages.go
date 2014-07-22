package models

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/russross/blackfriday"
)

type Page struct {
	Id       int
	Title    string
	User     string
	Locked   bool
	Created  time.Time
	Modified mysql.NullTime
	Content  sql.NullString
}

func (this *Page) Load(title string) error {
	row := DB.QueryRow("SELECT id, title, user, locked, created, modified, content FROM pages WHERE title = ?", title)
	err := row.Scan(&this.Id, &this.Title, &this.User, &this.Locked, &this.Created, &this.Modified, &this.Content)
	return err
}
func (this *Page) Save(title string) error {

	if title == "" || !PageExists(title) {
		_, err := DB.Exec("INSERT INTO pages ( title, user, locked, created, modified, content ) VALUES ( ?, ?, ?, ?, ?, ?)", this.Title, this.User, this.Locked, this.Created, this.Modified, this.Content)
		if err != nil {
			return err
		}
	} else {
		_, err := DB.Exec("UPDATE pages SET title = ?, user = ?, locked = ?, created = ?, modified = ?, content = ? WHERE title = ?", this.Title, this.User, this.Locked, this.Created, this.Modified, this.Content, title)
		if err != nil {
			return err
		}
	}

	return nil
}
func (this *Page) Markdown() []byte {
	return blackfriday.MarkdownCommon([]byte(this.Content.String))
}

func GetPageList(pages *[]Page, limit int) error {

	var rowCount int

	err := DB.QueryRow("SELECT count(id) FROM pages").Scan(&rowCount)
	if err != nil {
		return err
	}

	if cap(*pages) != rowCount {
		*pages = make([]Page, rowCount)
	}

	rows, err := DB.Query("SELECT id, title, user, locked, created, modified, content FROM pages")
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		err := rows.Scan(&(*pages)[i].Id, &(*pages)[i].Title, &(*pages)[i].User, &(*pages)[i].Locked, &(*pages)[i].Created, &(*pages)[i].Modified, &(*pages)[i].Content)
		if err != nil {
			return err
		}

		i++
		if i >= limit {
			break
		}
	}

	return nil
}

func PageExists(title string) bool {

	var rowCount int

	err := DB.QueryRow("SELECT count(id) FROM pages WHERE title=?", title).Scan(&rowCount)
	if err != nil {
		return false
	}

	return rowCount > 0
}
