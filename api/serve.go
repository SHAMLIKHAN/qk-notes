package api

import (
	"log"
	"net/http"
)

// Serve : To run api server
func (a *App) Serve(addr string) {
	router := a.Router()
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, router)
}
