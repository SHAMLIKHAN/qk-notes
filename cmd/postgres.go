package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"qk-note/shared"
)

func preparePostgres() (*sql.DB, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDatabaseURL() (string, error) {
	env, err := getEnv()
	if err != nil {
		return shared.EmptyString, err
	}
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env[shared.DBHOST],
		env[shared.DBPORT],
		env[shared.DBUSER],
		env[shared.DBPASSWORD],
		env[shared.DBNAME],
	)
	return psql, nil
}

func getEnv() (map[string]string, error) {
	env := make(map[string]string)
	host, ok := os.LookupEnv(shared.HOST)
	if !ok {
		return nil, errors.New("HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(shared.PORT)
	if !ok {
		return nil, errors.New("PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(shared.USER)
	if !ok {
		return nil, errors.New("USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(shared.PASSWORD)
	if !ok {
		return nil, errors.New("PASSWORD environment variable required but not set")
	}
	database, ok := os.LookupEnv(shared.DATABASE)
	if !ok {
		return nil, errors.New("DATABASE environment variable required but not set")
	}
	env[shared.DBHOST] = host
	env[shared.DBPORT] = port
	env[shared.DBUSER] = user
	env[shared.DBPASSWORD] = password
	env[shared.DBNAME] = database
	return env, nil
}
