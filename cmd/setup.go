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
		env[consts.DBHOST],
		env[consts.DBPORT],
		env[consts.DBUSER],
		env[consts.DBPASSWORD],
		env[consts.DBNAME],
	)
	return psql, nil
}

func getEnv() (map[string]string, error) {
	env := make(map[string]string)
	host, ok := os.LookupEnv(consts.HOST)
	if !ok {
		return nil, errors.New("HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(consts.PORT)
	if !ok {
		return nil, errors.New("PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(consts.USER)
	if !ok {
		return nil, errors.New("USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(consts.PASSWORD)
	if !ok {
		return nil, errors.New("PASSWORD environment variable required but not set")
	}
	database, ok := os.LookupEnv(consts.DATABASE)
	if !ok {
		return nil, errors.New("DATABASE environment variable required but not set")
	}
	env[consts.DBHOST] = host
	env[consts.DBPORT] = port
	env[consts.DBUSER] = user
	env[consts.DBPASSWORD] = password
	env[consts.DBNAME] = database
	return env, nil
}
