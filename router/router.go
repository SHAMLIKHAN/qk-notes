package router

import (
	"database/sql"
	"qk-note/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router : chi Router
func Router(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	uah := user.NewHTTPHandler(db)

	r.Post("/register", uah.RegisterUser)

	return r
}
