package note

import (
	"database/sql"
)

// PostgresRepo : Note Repo Struct for Postgres
type PostgresRepo struct {
	DB *sql.DB
}

// Create : Postgres function to Create a Note
func (pg *PostgresRepo) Create(note *Note) error {
	query := `INSERT INTO note (user_id, heading, content, status, category, tags) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := pg.DB.Exec(query, note.UserID, note.Heading, note.Content, note.Status, note.Category, note.Tags)
	return err
}
