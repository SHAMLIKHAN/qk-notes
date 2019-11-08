package api

import (
	"log"
	"net/http"
)

// Serve : Run api server
func (a *App) Serve(addr string) {
	router := a.Router()
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, router)
}
