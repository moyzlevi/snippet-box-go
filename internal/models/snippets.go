package models

import (
	"database/sql"
	"time"
	"errors"
)

type Snippet struct {
	id int
	title string
	snippet_content string
	created_on time.Time
	expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippet.snippets  (title,snippet_content,created_on,expires)
	VALUES(
		$1, 
		$2,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP + interval '1 day' * $3) RETURNING id`
	
	var snippetId int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&snippetId)
	
	if err != nil {
		return 0, err
	}

	return snippetId, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, snippet_content, created_on, expires from snippet.snippets
	WHERE expires > CURRENT_TIMESTAMP AND id = $1`
	
	s := &Snippet{}

	err := m.DB.QueryRow(stmt, id).Scan(&s.id,
		&s.title,&s.snippet_content,&s.created_on,&s.expires)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest(id int) ([]*Snippet, error) {
	return nil, nil
}