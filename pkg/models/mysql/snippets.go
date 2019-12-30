package mysql

import (
	"database/sql"

	"github.com/rajaseelan/snippetbox/pkg/models"
)

// SnippetModel Define a snippetModel type wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// SQL statement we want to execute
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use Exec() method
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use LastInsertId() method on result object to get the ID of our
	// newly inserted record in the snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// ID returned will be int64, convert to int before return
	return int(id), nil
}

// Get return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
