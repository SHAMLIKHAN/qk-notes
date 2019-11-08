package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"qk-note/consts"

	// go package for postgres
	_ "github.com/lib/pq"
)

const (
	// HOST : Environment variable for database host
	HOST = "QKNOTE_HOST"
	// PORT : Environment variable for database port
	PORT = "QKNOTE_PORT"
	// USER : Environment variable for database user
	USER = "QKNOTE_USER"
	// PASSWORD : Environment variable for database password
	PASSWORD = "QKNOTE_PASSWORD"
	// DATABASE : Environment variable for database name
	DATABASE = "QKNOTE_DATABASE"
	// DBHOST : Database host
	DBHOST = "HOST"
	// DBPORT : Database port
	DBPORT = "PORT"
	// DBUSER : Database user
	DBUSER = "USER"
	// DBPASSWORD : Database password
	DBPASSWORD = "PASSWORD"
	// DBNAME : Database name
	DBNAME = "DATABASE"
)

func prepareDatabase() (*sql.DB, error) {
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
	log.Println("App : Database connected successfully!")
	return db, nil
}

func getDatabaseURL() (string, error) {
	env, err := getEnv()
	if err != nil {
		return consts.EmptyString, err
	}
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env[DBHOST],
		env[DBPORT],
		env[DBUSER],
		env[DBPASSWORD],
		env[DBNAME],
	)
	return psql, nil
}

func getEnv() (map[string]string, error) {
	env := make(map[string]string)
	host, ok := os.LookupEnv(HOST)
	if !ok {
		return nil, errors.New("HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(PORT)
	if !ok {
		return nil, errors.New("PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(USER)
	if !ok {
		return nil, errors.New("USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(PASSWORD)
	if !ok {
		return nil, errors.New("PASSWORD environment variable required but not set")
	}
	database, ok := os.LookupEnv(DATABASE)
	if !ok {
		return nil, errors.New("DATABASE environment variable required but not set")
	}
	env[DBHOST] = host
	env[DBPORT] = port
	env[DBUSER] = user
	env[DBPASSWORD] = password
	env[DBNAME] = database
	return env, nil
}
