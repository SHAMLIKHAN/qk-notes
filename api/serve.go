package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

// Serve : Run api server
func (a *App) Serve(addr string) {
	router := a.Router()
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
}
