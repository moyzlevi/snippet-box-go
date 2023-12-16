package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID              int
	Title           string
	Snippet_content string
	Created_on      time.Time
	Expires         time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets  (title,snippet_content,created_on,expires)
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
	stmt := `SELECT id, title, snippet_content, created_on, expires from snippets
	WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	s := &Snippet{}

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID,
		&s.Title, &s.Snippet_content, &s.Created_on, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, snippet_content, created_on, expires from snippets
	WHERE expires > CURRENT_TIMESTAMP
	ORDER BY id 
	LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID,
			&s.Title, &s.Snippet_content, &s.Created_on, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
