package user

import "database/sql"

// Repo : User Repository
type Repo interface {
	Register(*User) error
	CheckUsername(user *User) (int, error)
	CheckEmail(user *User) (int, error)
	LoginUser(user *Login) (*User, error)
}

// NewRepo : Returns User Repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
