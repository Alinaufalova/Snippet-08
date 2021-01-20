package postgres
import (
	"github.com/jackc/pgx/v4/pgxpool"
	"snippet08/pkg/models"
)
// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	Pool * pgxpool.Pool
}
// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES(?, ?, CURRENT_TIMESTAMP, DATE_ADD(CURRENT_TIMESTAMP(), INTERVAL '? DAY'))`

	result, err := m.Pool.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

	result, err = m.Pool.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > CURRENT_TIMESTAMP AND id = ?`

	row := m.Pool.QueryRow(stmt, id)
	s := &models.Snippet{}

}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.Pool.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
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