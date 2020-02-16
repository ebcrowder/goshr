package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/ebcrowder/goshr/schema"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func ConnectSqlite() (*Sqlite, error) {
	os.Remove("../../../../sqlite-data/goshr.db")

	db, err := sql.Open("sqlite3", "../../../../sqlite-data/goshr.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Sqlite{db}, nil
}

func (p *Sqlite) Close() {
	p.DB.Close()
}

func (p *Sqlite) Insert(file *schema.File) (int, error) {
	query := `
		INSERT INTO file (id, name, key)
		VALUES (nextval('todo_id'), $1, $2)
		RETURNING id;
	`

	rows, err := p.DB.Query(query, file.Name, file.Key)
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (p *Sqlite) Delete(id int) error {
	query := `
		DELETE FROM file 
		WHERE id = $1;
	`

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (p *Sqlite) GetFiles() ([]schema.File, error) {
	query := `
		SELECT *
		FROM file 
		ORDER BY id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var fileList []schema.File
	for rows.Next() {
		var t schema.File
		if err := rows.Scan(&t.ID, &t.Name, &t.Key); err != nil {
			return nil, err
		}
		fileList = append(fileList, t)
	}

	return fileList, nil
}
