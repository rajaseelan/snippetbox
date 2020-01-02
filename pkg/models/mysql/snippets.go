package mysql

import (
	"database/sql"
	"errors"
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

	// SQL Statement we want to execute
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// use QueryRow() on the connection pool
	// returns a pointer to sql.Row object
	row := m.DB.QueryRow(stmt, id)

	// Initialize snippet to a new zeroed Snippet Struct
	s := &models.Snippet{}

	// use row.Scan() to copy values from each field to the corresponding snippet struct field
	// args to row.Scan are pointers, no of clumns must be the same
	// as no of columns returned
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	// everything went ok, return the snippet object
	return s, nil
}

// Latest return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
