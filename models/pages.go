package models

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Page struct {
	Id          int
	Title       string
	User        string
	Locked      bool
	Created     time.Time
	Modified    mysql.NullTime
	Content     sql.NullString
	Attachments Attachments
}

type Attachments struct {
	Files []File
}

type File struct {
	Name string
	Type string
	Data []byte
}

func (this *Page) LoadFromTitle(title string) error {

	id, err := GetPageIDFromTitle(title)
	if err != nil {
		return err
	}

	return this.Load(id)
}

func (this *Page) Load(id int) error {

	var attachments []byte

	row := DB.QueryRow("SELECT id, title, user, locked, created, modified, content, attachments FROM pages WHERE id = ?", id)
	err := row.Scan(&this.Id, &this.Title, &this.User, &this.Locked, &this.Created, &this.Modified, &this.Content, &attachments)

	if err != nil {
		return err
	}

	if attachments == nil {
		return nil
	}

	buffer := bytes.NewBuffer(attachments)
	decoder := gob.NewDecoder(buffer)

	err = decoder.Decode(&this.Attachments)

	return err

}

func (this *Page) SaveToTitle(title string) error {

	// 上書きの場合
	if PageExists(title) {
		id, err := GetPageIDFromTitle(title)
		if err != nil {
			return err
		}

		return this.Update(id)
	}

	// 新規作成の場合
	return this.Create()
}

func (this *Page) Create() error {

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(this.Attachments)
	if err != nil {
		return err
	}

	_, err = DB.Exec("INSERT INTO pages ( title, user, locked, created, modified, content, attachments ) VALUES ( ?, ?, ?, ?, ?, ?, ?)", this.Title, this.User, this.Locked, this.Created, this.Modified, this.Content, buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (this *Page) Update(id int) error {

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(this.Attachments)
	if err != nil {
		return err
	}

	_, err = DB.Exec("UPDATE pages SET title = ?, user = ?, locked = ?, created = ?, modified = ?, content = ?, attachments = ? WHERE id = ?", this.Title, this.User, this.Locked, this.Created, this.Modified, this.Content, buffer.Bytes(), id)
	if err != nil {
		return err
	}

	return nil
}

func (this *Page) Markdown() []byte {
	return bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(this.Content.String)))
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

func GetPageIDFromTitle(title string) (int, error) {

	var id int

	err := DB.QueryRow("SELECT id FROM pages WHERE title=?", title).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
