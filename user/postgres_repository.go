package user

import (
	"database/sql"
)

// PostgresRepo : User Repo Struct for Postgres
type PostgresRepo struct {
	DB *sql.DB
}

// Register : Postgres function to Register a User
func (pg *PostgresRepo) Register(user *User) error {
	query := `INSERT INTO user_account (firstname, lastname, username, email, password, role) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := pg.DB.Exec(query, user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Role)
	return err
}

// CheckUsername : Postgres function to check the username already exists or not
func (pg *PostgresRepo) CheckUsername(user *User) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM user_account WHERE username = $1;`
	err := pg.DB.QueryRow(query, user.Username).Scan(&count)
	return count, err
}

// CheckEmail : Postgres function to check the email already exists or not
func (pg *PostgresRepo) CheckEmail(user *User) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM user_account WHERE email = $1;`
	err := pg.DB.QueryRow(query, user.Email).Scan(&count)
	return count, err
}

// LoginUser : Postgres function to validate user crederntials
func (pg *PostgresRepo) LoginUser(login *Login) (*User, error) {
	var user User
	query := `SELECT * FROM user_account WHERE (username = $1 OR email = $2) AND password = $3 AND status = $4;`
	row := pg.DB.QueryRow(query, login.Username, login.Email, login.Password, ALIVE)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Role, &user.Status)
	return &user, err
}
