package note

import (
	"database/sql"
)

// PostgresRepo : Note Repo Struct for Postgres
type PostgresRepo struct {
	DB *sql.DB
}

// Create : Postgres function to Create a Note
func (pg *PostgresRepo) Create(note *Note) (*Note, error) {
	var n Note
	query := `INSERT INTO note (user_id, heading, content, status, category, tags) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;`
	row := pg.DB.QueryRow(query, note.UserID, note.Heading, note.Content, note.Status, note.Category, note.Tags)
	err := row.Scan(&n.ID, &n.UserID, &n.Heading, &n.Content, &n.Status, &n.Category, &n.Tags)
	return &n, err
}
