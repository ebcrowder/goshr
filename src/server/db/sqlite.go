package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type File struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Sqlite struct {
	DB *sql.DB
}

func ConnectSqlite() (*Sqlite, error) {
	var connStr = "~/sqlite-data/goshr.db"
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Sqlite{db}, nil
}

func (p *Sqlite) Close() {
	p.DB.Close()
}

func (p *Sqlite) Insert(file *File) (int, error) {
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

func (p *Sqlite) GetFiles() ([]File, error) {
	query := `
		SELECT *
		FROM file 
		ORDER BY id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var fileList []File
	for rows.Next() {
		var t File
		if err := rows.Scan(&t.ID, &t.Name, &t.Key); err != nil {
			return nil, err
		}
		fileList = append(fileList, t)
	}

	return fileList, nil
}
