package router

import (
	"database/sql"
	"qk-note/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router : Basic Router
type Router interface {
	Setup() *chi.Mux
}

// ChiRouter : Struct that holds router and DB connection
type ChiRouter struct {
	DB *sql.DB
}

// NewRouter : Returns Basic Router
func NewRouter(db *sql.DB) Router {
	return &ChiRouter{
		DB: db,
	}
}

// Setup : chi Router
func (r *ChiRouter) Setup() *chi.Mux {
	cr := chi.NewRouter()
	cr.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	uah := user.NewHTTPHandler(r.DB)

	cr.Route("/user", func(cr chi.Router) {
		cr.Post("/register", uah.RegisterUser)
	})

	return cr
}
