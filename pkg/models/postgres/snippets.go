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
	return 0, nil
}
// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}
// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}