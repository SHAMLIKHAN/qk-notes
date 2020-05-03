package note

import "database/sql"

// Repo : Note Repository
type Repo interface {
	Create(*Note) error
}

// NewRepo : Returns Note Repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
