package core

import (
	"encoding/json"
	"log"
	"net/http"
	"qk-note/consts"
	"qk-note/models"
	"qk-note/shared"
)

// CreateNote : to add a new note
func (c *Capsule) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	body := json.NewDecoder(r.Body)
	err := body.Decode(&note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 400, consts.DecodeErrorCode, consts.DecodeError)
		return
	}
	note.UserID = 1
	note.Status = 1
	err = c.insertNoteIntoDatabase(note)
	if err != nil {
		log.Println("App : Error! ", err.Error())
		shared.Fail(w, 500, consts.DatabaseErrorCode, consts.DatabaseError)
		return
	}
	shared.Send(w, 200, note)
}

// GetNotes : to return all notes
func (c *Capsule) GetNotes(w http.ResponseWriter, r *http.Request) {
	shared.Send(w, 200, "Success")
}

// GetNote : to get a particular note
func (c *Capsule) GetNote(w http.ResponseWriter, r *http.Request) {
	shared.Send(w, 200, "Success")
}

// EditNote : to update a particular note
func (c *Capsule) EditNote(w http.ResponseWriter, r *http.Request) {
	shared.Send(w, 200, "Success")
}

// DeleteNote : to delete a particular note
func (c *Capsule) DeleteNote(w http.ResponseWriter, r *http.Request) {
	shared.Send(w, 200, "Success")
}
