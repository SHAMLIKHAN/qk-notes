package core

import "net/http"

// QK : An interface to capsule
type QK interface {
	CreateNote(http.ResponseWriter, *http.Request)
	GetNotes(http.ResponseWriter, *http.Request)
	GetNote(http.ResponseWriter, *http.Request)
	EditNote(http.ResponseWriter, *http.Request)
	DeleteNote(http.ResponseWriter, *http.Request)
}
