package api

import "github.com/gorilla/mux"

// Router : Basic router
func (a *App) Router() *mux.Router {
	r := mux.NewRouter()
	noteRouter := r.PathPrefix("/note").Subrouter()

	noteRouter.HandleFunc("", a.CreateNote).Methods("POST")
	noteRouter.HandleFunc("", a.GetNotes).Methods("GET")
	noteRouter.HandleFunc("/{id}", a.GetNote).Methods("GET")
	noteRouter.HandleFunc("/{id}", a.EditNote).Methods("PATCH")
	noteRouter.HandleFunc("/{id}", a.DeleteNote).Methods("DELETE")

	return r
}
