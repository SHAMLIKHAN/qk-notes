package cmd

import (
	"database/sql"
	"log"
	"net/http"
	"qk-note/router"
	"qk-note/user"
)

// App : Struct to represent this app
type App struct {
	user.Handler
}

// NewApp : to get App Struct
func NewApp(db *sql.DB) *App {
	return &App{
		Handler: user.NewHTTPHandler(db),
	}
}

// Serve : to Run API Server
func (a *App) Serve(addr string, db *sql.DB) {
	router := router.Router(db)
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, router)
}
