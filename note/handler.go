package note

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"qk-note/shared"
	"strconv"
)

// Handler : Note Handler
type Handler interface {
	CreateNote(w http.ResponseWriter, r *http.Request)
}

// HandleNote : Note Handler Struct
type HandleNote struct {
	ns Service
}

// NewHTTPHandler : Returns Note HTTP Handler
func NewHTTPHandler(db *sql.DB) Handler {
	return &HandleNote{
		ns: NewNoteService(db),
	}
}

// CreateNote : to Create a Note
func (nh *HandleNote) CreateNote(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /user/note API hit!")
	var note Note
	body := json.NewDecoder(r.Body)
	err := body.Decode(&note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.DecodeErrorCode, shared.DecodeError)
		return
	}
	id := r.Header.Get("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.DecodeErrorCode, shared.DecodeError)
		return
	}
	err = nh.ns.ValidateNote(&note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, shared.InputDataErrorCode, err.Error())
		return
	}
	note.UserID = userID
	note.Status = ACTIVE
	note.Category = DEFAULT
	note.Tags = ""
	err = nh.ns.CreateNote(&note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, shared.DatabaseErrorCode, shared.DatabaseError)
		return
	}
	log.Println("App : Note created successfully! ", note)
	shared.Send(w, 200, note)
	return
}
