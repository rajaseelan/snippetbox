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
	return 0, nil
}

// Get return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
